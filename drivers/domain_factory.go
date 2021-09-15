package databases

import (
	userDomain "peduli-covid/businesses/users"
	userDB "peduli-covid/drivers/postgres/users"

	provinceDomain "peduli-covid/businesses/provinces"
	provinceDB "peduli-covid/drivers/postgres/provinces"

	roleDomain "peduli-covid/businesses/roles"
	roleDB "peduli-covid/drivers/postgres/roles"

	cityDomain "peduli-covid/businesses/cities"
	cityDB "peduli-covid/drivers/postgres/cities"

	hospitalDomain "peduli-covid/businesses/hospitals"
	hospitalDB "peduli-covid/drivers/postgres/hospitals"

	bedtypeDomain "peduli-covid/businesses/bedtypes"
	bedtypeDB "peduli-covid/drivers/postgres/bedtypes"

	reservationDomain "peduli-covid/businesses/reservations"
	reservationDB "peduli-covid/drivers/postgres/reservations"

	invoiceDomain "peduli-covid/businesses/invoices"
	invoiceDB "peduli-covid/drivers/postgres/invoices"

	paymentDomain "peduli-covid/businesses/payments"
	paymentDB "peduli-covid/drivers/postgres/payments"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewPostgresRepository(conn)
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

func NewHospitalRepository(conn *gorm.DB) hospitalDomain.Repository {
	return hospitalDB.NewPostgresRepository(conn)
}

func NewBedtypeRepository(conn *gorm.DB) bedtypeDomain.Repository {
	return bedtypeDB.NewPostgresRepository(conn)
}

func NewReservationRepository(conn *gorm.DB) reservationDomain.Repository {
	return reservationDB.NewPostgresRepository(conn)
}

func NewInvoiceRepository(conn *gorm.DB) invoiceDomain.Repository {
	return invoiceDB.NewPostgresRepository(conn)
}

func NewPaymentRepository(conn *gorm.DB) paymentDomain.Repository {
	return paymentDB.NewPostgresRepository(conn)
}
