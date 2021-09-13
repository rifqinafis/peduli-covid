package databases

import (
	userDomain "peduli-covid/businesses/users"
	userDB "peduli-covid/drivers/postgres/users"

	adminDomain "peduli-covid/businesses/admins"
	adminDB "peduli-covid/drivers/postgres/admins"

	provinceDomain "peduli-covid/businesses/provinces"
	provinceDB "peduli-covid/drivers/postgres/provinces"

	roleDomain "peduli-covid/businesses/roles"
	roleDB "peduli-covid/drivers/postgres/roles"

	cityDomain "peduli-covid/businesses/cities"
	cityDB "peduli-covid/drivers/postgres/cities"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgresRepository(conn)
}

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewPostgresRepository(conn)
}

func NewProvinceRepository(conn *gorm.DB) provinceDomain.Repository {
	return provinceDB.NewPostgresRepository(conn)
}

func NewRoleRepository(conn *gorm.DB) roleDomain.Repository {
	return roleDB.NewPostgresRepository(conn)
}

func NewCityRepository(conn *gorm.DB) cityDomain.Repository {
	return cityDB.NewPostgresRepository(conn)
}
