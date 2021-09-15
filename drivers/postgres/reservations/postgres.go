package reservations

import (
	"context"
	"peduli-covid/businesses/reservations"

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

func (nr *PostgresRepository) FindAll(ctx context.Context) ([]reservations.Domain, error) {
	rec := []Reservations{}

	err := nr.Conn.Find(&rec).Error
	if err != nil {
		return []reservations.Domain{}, err
	}

	var domainReservations []reservations.Domain
	for _, value := range rec {
		domainReservations = append(domainReservations, value.toDomain())
	}
	return domainReservations, nil
}

func (nr *PostgresRepository) FindByUserID(ctx context.Context, userID int) ([]reservations.Domain, error) {
	rec := []Reservations{}

	err := nr.Conn.Where("user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []reservations.Domain{}, err
	}

	var domainReservations []reservations.Domain
	for _, value := range rec {
		domainReservations = append(domainReservations, value.toDomain())
	}
	return domainReservations, nil
}

func (nr *PostgresRepository) FindByHospitalID(ctx context.Context, hospitalID int) ([]reservations.Domain, error) {
	rec := []Reservations{}

	err := nr.Conn.Where("hospital_id = ?", hospitalID).Find(&rec).Error
	if err != nil {
		return []reservations.Domain{}, err
	}

	var domainReservations []reservations.Domain
	for _, value := range rec {
		domainReservations = append(domainReservations, value.toDomain())
	}
	return domainReservations, nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, id int) (reservations.Domain, error) {
	rec := Reservations{}
	err := nr.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return reservations.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) UpdateStatus(ctx context.Context, reservationDomain *reservations.Domain) error {
	rec := fromDomain(*reservationDomain)

	result := nr.Conn.Table("reservations").Where("id = ?", rec.ID).Update("status", &rec.Status)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *PostgresRepository) Store(ctx context.Context, reservationDomain *reservations.Domain) (reservations.Domain, error) {
	rec := fromDomain(*reservationDomain)
	resp := Reservations{}
	err := nr.Conn.Create(rec).Scan(&resp).Error
	if err != nil {
		return reservations.Domain{}, err
	}

	return resp.toDomain(), nil
}
