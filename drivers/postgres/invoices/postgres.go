package invoices

import (
	"context"
	"peduli-covid/businesses/invoices"

	"gorm.io/gorm"
)

type PostgresRepository struct {
	Conn *gorm.DB
}

func NewPostgresRepository(conn *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		Conn: conn,
	}
}

func (nr *PostgresRepository) FindByUserID(ctx context.Context, userID int) ([]invoices.Domain, error) {
	rec := []Invoices{}

	err := nr.Conn.Joins("left join reservations on reservations.id = invoices.reservation_id").Where("reservations.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []invoices.Domain{}, err
	}

	var domainInvoices []invoices.Domain
	for _, value := range rec {
		domainInvoices = append(domainInvoices, value.toDomain())
	}
	return domainInvoices, nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, id int) (invoices.Domain, error) {
	rec := Invoices{}
	err := nr.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return invoices.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, invoiceDomain *invoices.Domain) (invoices.Domain, error) {
	rec := fromDomain(*invoiceDomain)
	resp := Invoices{}
	err := nr.Conn.Create(rec).Scan(&resp).Error
	if err != nil {
		return invoices.Domain{}, err
	}

	return resp.toDomain(), nil
}
