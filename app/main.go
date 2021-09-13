package main

import (
	_dbFactory "peduli-covid/drivers"

	_userUsecase "peduli-covid/businesses/users"
	_userController "peduli-covid/controllers/users"
	_userRepo "peduli-covid/drivers/postgres/users"

	_adminUsecase "peduli-covid/businesses/admins"
	_adminController "peduli-covid/controllers/admins"
	_adminRepo "peduli-covid/drivers/postgres/admins"

	_rsbedcovidUsecase "peduli-covid/businesses/rsbedcovid"
	_rsbedcovidController "peduli-covid/controllers/rsbedcovid"
	_rsbedcovidRepo "peduli-covid/drivers/thirdparties/rsbedcovid"

	_provinceUsecase "peduli-covid/businesses/provinces"
	_provinceController "peduli-covid/controllers/provinces"
	_provinceRepo "peduli-covid/drivers/postgres/provinces"

	_roleUsecase "peduli-covid/businesses/roles"
	_roleController "peduli-covid/controllers/roles"
	_roleRepo "peduli-covid/drivers/postgres/roles"

	_cityUsecase "peduli-covid/businesses/cities"
	_cityController "peduli-covid/controllers/cities"
	_cityRepo "peduli-covid/drivers/postgres/cities"

	_dbDriver "peduli-covid/drivers/postgres"

	_config "peduli-covid/app/config"
	_middleware "peduli-covid/app/middleware"
	_routes "peduli-covid/app/routes"

	"log"
	"time"

	echo "github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_userRepo.Users{},
		&_adminRepo.Admins{},
		&_provinceRepo.Provinces{},
		&_roleRepo.Roles{},
		&_cityRepo.Cities{},
	)
}

func main() {
	configApp := _config.GetConfig()
	configDB := _dbDriver.ConfigDB{
		DB_Username: configApp.Database.User,
		DB_Password: configApp.Database.Pass,
		DB_Host:     configApp.Database.Host,
		DB_Port:     configApp.Database.Port,
		DB_Database: configApp.Database.Name,
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	e := echo.New()

	userRepo := _dbFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	adminRepo := _dbFactory.NewAdminRepository(db)
	adminUsecase := _adminUsecase.NewAdminUsecase(adminRepo, &configJWT, timeoutContext)
	adminCtrl := _adminController.NewAdminController(adminUsecase)

	rsbedcovidRepo := _rsbedcovidRepo.NewRSBedCovid()
	rsbedcovidUsecase := _rsbedcovidUsecase.NewRSBedCovid(rsbedcovidRepo, &configJWT, timeoutContext)
	rsbedcovidCtrl := _rsbedcovidController.NewRSBedCovidController(rsbedcovidUsecase)

	provinceRepo := _dbFactory.NewProvinceRepository(db)
	provinceUsecase := _provinceUsecase.NewProvinceUsecase(provinceRepo, rsbedcovidRepo, &configJWT, timeoutContext)
	provinceCtrl := _provinceController.NewProvinceController(provinceUsecase)

	roleRepo := _dbFactory.NewRoleRepository(db)
	roleUsecase := _roleUsecase.NewRoleUsecase(roleRepo, &configJWT, timeoutContext)
	roleCtrl := _roleController.NewRoleController(roleUsecase)

	cityRepo := _dbFactory.NewCityRepository(db)
	cityUsecase := _cityUsecase.NewCityUsecase(cityRepo, rsbedcovidRepo, &configJWT, timeoutContext)
	cityCtrl := _cityController.NewCityController(cityUsecase)

	routesInit := _routes.ControllerList{
		JWTMiddleware:        configJWT.Init(),
		AdminController:      *adminCtrl,
		UserController:       *userCtrl,
		RSBedCovidController: *rsbedcovidCtrl,
		ProvinceController:   *provinceCtrl,
		RoleController:       *roleCtrl,
		CityController:       *cityCtrl,
	}
	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
