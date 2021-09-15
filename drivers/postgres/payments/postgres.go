package payments

import (
	"context"
	"peduli-covid/businesses/payments"

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

func (nr *PostgresRepository) FindByUserID(ctx context.Context, userID int) ([]payments.Domain, error) {
	rec := []Payments{}

	err := nr.Conn.Joins("left join reservations on reservations.id = payments.reservation_id").Where("reservations.user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []payments.Domain{}, err
	}

	var domainPayments []payments.Domain
	for _, value := range rec {
		domainPayments = append(domainPayments, value.toDomain())
	}
	return domainPayments, nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, id int) (payments.Domain, error) {
	rec := Payments{}
	err := nr.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return payments.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, invoiceDomain *payments.Domain) error {
	rec := fromDomain(*invoiceDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
