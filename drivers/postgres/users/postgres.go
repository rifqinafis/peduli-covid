package users

import (
	"context"
	"peduli-covid/businesses/users"

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

func (nr *PostgresRepository) FindByHospitalID(ctx context.Context, hospitalID int) ([]users.Domain, error) {
	rec := []Users{}

	err := nr.Conn.Where("hospital_id = ?", hospitalID).Find(&rec).Error
	if err != nil {
		return []users.Domain{}, err
	}

	var domainUsers []users.Domain
	for _, value := range rec {
		domainUsers = append(domainUsers, value.toDomain())
	}
	return domainUsers, nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, userID int) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("id = ?", userID).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) GetByEmail(ctx context.Context, email string) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return users.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, userDomain *users.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
