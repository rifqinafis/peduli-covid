package databases

import (
	userDomain "peduli-covid/businesses/users"
	userDB "peduli-covid/drivers/databases/users"

	adminDomain "peduli-covid/businesses/admins"
	adminDB "peduli-covid/drivers/databases/admins"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgresRepository(conn)
}

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewPostgresRepository(conn)
}
