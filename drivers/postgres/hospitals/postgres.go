package hospitals

import (
	"context"
	"peduli-covid/businesses/hospitals"

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

func (nr *PostgresRepository) FindAll(ctx context.Context) ([]hospitals.Domain, error) {
	rec := []Hospitals{}

	err := nr.Conn.Find(&rec).Error
	if err != nil {
		return []hospitals.Domain{}, err
	}

	var domainHospitals []hospitals.Domain
	for _, value := range rec {
		domainHospitals = append(domainHospitals, value.toDomain())
	}
	return domainHospitals, nil
}

func (nr *PostgresRepository) FindByCityID(ctx context.Context, cityID int) ([]hospitals.Domain, error) {
	rec := []Hospitals{}

	err := nr.Conn.Where("city_id = ?", cityID).Find(&rec).Error
	if err != nil {
		return []hospitals.Domain{}, err
	}

	var domainHospitals []hospitals.Domain
	for _, value := range rec {
		domainHospitals = append(domainHospitals, value.toDomain())
	}
	return domainHospitals, nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, id int) (hospitals.Domain, error) {
	rec := Hospitals{}
	err := nr.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return hospitals.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, userDomain *hospitals.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
