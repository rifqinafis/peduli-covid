package provinces

import (
	"context"
	"peduli-covid/businesses/provinces"

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

func (nr *PostgresRepository) FindAll(ctx context.Context) ([]provinces.Domain, error) {
	rec := []Provinces{}

	err := nr.Conn.Find(&rec).Error
	if err != nil {
		return []provinces.Domain{}, err
	}

	var domainProvinces []provinces.Domain
	for _, value := range rec {
		domainProvinces = append(domainProvinces, value.toDomain())
	}
	return domainProvinces, nil
}

func (nr *PostgresRepository) GetByCode(ctx context.Context, code string) (provinces.Domain, error) {
	rec := Provinces{}
	err := nr.Conn.Where("code = ?", code).First(&rec).Error
	if err != nil {
		return provinces.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, userDomain *provinces.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
