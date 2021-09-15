package routes

import (
	"errors"
	"net/http"
	"peduli-covid/app/middleware"
	controller "peduli-covid/controllers"
	"peduli-covid/controllers/admins"
	"peduli-covid/controllers/bedtypes"
	"peduli-covid/controllers/cities"
	"peduli-covid/controllers/hospitals"
	"peduli-covid/controllers/invoices"
	"peduli-covid/controllers/payments"
	"peduli-covid/controllers/provinces"
	"peduli-covid/controllers/reservations"
	"peduli-covid/controllers/roles"
	rsbedcovids "peduli-covid/controllers/rsbedcovid"
	"peduli-covid/controllers/users"
	consts "peduli-covid/helpers/const"

	echo "github.com/labstack/echo/v4"
	midware "github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware         midware.JWTConfig
	AdminController       admins.AdminController
	UserController        users.UserController
	RSBedCovidController  rsbedcovids.RSBedCovidController
	ProvinceController    provinces.ProvinceController
	RoleController        roles.RoleController
	CityController        cities.CityController
	HospitalController    hospitals.HospitalController
	BedtypeController     bedtypes.BedtypeController
	ReservationController reservations.ReservationController
	InvoiceController     invoices.InvoiceController
	PaymentController     payments.PaymentController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", cl.UserController.Store)
	users.POST("/login", cl.UserController.Login)

	admins := e.Group("admins")
	admins.POST("/register", cl.AdminController.Store)
	admins.POST("/login", cl.AdminController.Login)

	rsbedcovid := e.Group("/rsbedcovid", midware.JWTWithConfig(cl.JWTMiddleware))
	rsbedcovid.GET("/province", cl.RSBedCovidController.GetProvince)
	rsbedcovid.GET("/city", cl.RSBedCovidController.GetCity)
	rsbedcovid.GET("/hospital", cl.RSBedCovidController.GetHospital)
	rsbedcovid.GET("/bed-detail", cl.RSBedCovidController.GetBedDetail)
	rsbedcovid.GET("/location", cl.RSBedCovidController.GetHospitalLocation)

	clone := e.Group("/clone", midware.JWTWithConfig(cl.JWTMiddleware))
	clone.POST("/province", cl.ProvinceController.StoreFromAPI)
	clone.POST("/city", cl.CityController.StoreFromAPI)
	clone.POST("/hospital", cl.HospitalController.StoreFromAPI)
	clone.POST("/bedtype", cl.BedtypeController.StoreFromAPI)

	province := e.Group("/province", midware.JWTWithConfig(cl.JWTMiddleware))
	province.GET("/", cl.ProvinceController.FindAll)

	city := e.Group("/city", midware.JWTWithConfig(cl.JWTMiddleware))
	city.GET("/", cl.CityController.FindByProvinceCode)

	hospital := e.Group("/hospital", midware.JWTWithConfig(cl.JWTMiddleware))
	hospital.GET("/", cl.HospitalController.FindByCityID)

	bedtype := e.Group("/bedtype", midware.JWTWithConfig(cl.JWTMiddleware))
	bedtype.GET("/", cl.BedtypeController.FindByHospitalID)

	reservation := e.Group("/reservation", midware.JWTWithConfig(cl.JWTMiddleware))
	reservation.POST("/", cl.ReservationController.Store, RoleValidation(consts.USER_ROLE, cl.UserController))
	reservation.PUT("/id/:id", cl.ReservationController.UpdateStatus, RoleValidation(consts.ADMIN_ROLE, cl.UserController))
	reservation.GET("/admin", cl.ReservationController.FindByAdminID, RoleValidation(consts.ADMIN_ROLE, cl.UserController))
	reservation.GET("/user", cl.ReservationController.FindByUserID, RoleValidation(consts.USER_ROLE, cl.UserController))

	invoice := e.Group("/invoice", midware.JWTWithConfig(cl.JWTMiddleware))
	invoice.GET("/id/:id", cl.InvoiceController.GetByID, RoleValidation(consts.USER_ROLE, cl.UserController))
	invoice.GET("/user", cl.InvoiceController.FindByUserID, RoleValidation(consts.USER_ROLE, cl.UserController))

	payment := e.Group("/payment", midware.JWTWithConfig(cl.JWTMiddleware))
	payment.POST("/", cl.PaymentController.Store, RoleValidation(consts.USER_ROLE, cl.UserController))
	payment.GET("/user", cl.PaymentController.FindByUserID, RoleValidation(consts.USER_ROLE, cl.UserController))
}

func RoleValidation(role string, userControler users.UserController) echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middleware.GetUser(c)
			userRole := userControler.UserRole(claims.ID)

			if userRole == role {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("forbidden roles"))
			}
		}
	}
}
