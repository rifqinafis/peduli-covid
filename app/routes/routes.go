package routes

import (
	"errors"
	"net/http"
	"peduli-covid/app/middleware"
	controller "peduli-covid/controllers"
	"peduli-covid/controllers/admins"
	"peduli-covid/controllers/cities"
	"peduli-covid/controllers/provinces"
	"peduli-covid/controllers/roles"
	rsbedcovids "peduli-covid/controllers/rsbedcovid"
	"peduli-covid/controllers/users"

	echo "github.com/labstack/echo/v4"
	midware "github.com/labstack/echo/v4/middleware"
)

const (
	adminCode = "admin"
	userCode  = "user"
)

type ControllerList struct {
	JWTMiddleware        midware.JWTConfig
	AdminController      admins.AdminController
	UserController       users.UserController
	RSBedCovidController rsbedcovids.RSBedCovidController
	ProvinceController   provinces.ProvinceController
	RoleController       roles.RoleController
	CityController       cities.CityController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.POST("/login", cl.UserController.Login)

	admins := e.Group("admins")
	admins.POST("/register", cl.AdminController.Store)
	admins.POST("/login", cl.AdminController.Login)

	rsbedcovid := e.Group("/rsbedcovid", midware.JWTWithConfig(cl.JWTMiddleware))
	rsbedcovid.GET("/province", cl.RSBedCovidController.GetProvince, RoleValidation(adminCode, cl.RoleController))
	rsbedcovid.GET("/city", cl.RSBedCovidController.GetCity, RoleValidation(adminCode, cl.RoleController))
	rsbedcovid.GET("/hospital", cl.RSBedCovidController.GetHospital, RoleValidation(adminCode, cl.RoleController))
	rsbedcovid.GET("/bed-detail", cl.RSBedCovidController.GetBedDetail, RoleValidation(adminCode, cl.RoleController))
	rsbedcovid.GET("/location", cl.RSBedCovidController.GetHospitalLocation, RoleValidation(adminCode, cl.RoleController))

	province := e.Group("/province", midware.JWTWithConfig(cl.JWTMiddleware))
	province.POST("/clone", cl.ProvinceController.StoreFromAPI)

	city := e.Group("/city", midware.JWTWithConfig(cl.JWTMiddleware))
	city.POST("/clone", cl.CityController.StoreFromAPI)
}

func RoleValidation(roleID string, roleController roles.RoleController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middleware.GetUser(c)
			userRole := roleController.FindRole(claims.ID)

			if userRole == roleID {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("invalid role"))
			}
		}
	}
}
