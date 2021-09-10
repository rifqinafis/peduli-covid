package routes

import (
	"peduli-covid/app/middleware"
	"peduli-covid/controllers/admins"
	rsbedcovids "peduli-covid/controllers/rsbedcovid"
	"peduli-covid/controllers/users"

	echo "github.com/labstack/echo/v4"
)

type ControllerList struct {
	JWTMiddleware        *middleware.ConfigJWT
	AdminController      admins.AdminController
	UserController       users.UserController
	RSBedCovidController rsbedcovids.RSBedCovidController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	adminMiddleware := cl.JWTMiddleware
	adminMiddleware.Role = []int{1}
	userMiddleware := cl.JWTMiddleware
	userMiddleware.Role = []int{2}

	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.POST("/login", cl.UserController.Login)

	admins := e.Group("admins")
	admins.POST("/register", cl.AdminController.Store)
	admins.POST("/login", cl.AdminController.Login)

	rsbedcovid := e.Group("/rsbedcovid", userMiddleware.VerifyRole)
	rsbedcovid.GET("/province", cl.RSBedCovidController.GetProvince)
	rsbedcovid.GET("/city", cl.RSBedCovidController.GetCity)
	rsbedcovid.GET("/hospital", cl.RSBedCovidController.GetHospital)
	rsbedcovid.GET("/bed-detail", cl.RSBedCovidController.GetBedDetail)
	rsbedcovid.GET("/location", cl.RSBedCovidController.GetHospitalLocation)
}
