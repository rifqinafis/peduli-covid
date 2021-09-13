package cities

import (
	"context"
	"peduli-covid/businesses/cities"

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

func (nr *PostgresRepository) Fetch(ctx context.Context, page, perpage int) ([]cities.Domain, int, error) {
	rec := []Cities{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []cities.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []cities.Domain{}, 0, err
	}

	var domainCities []cities.Domain
	for _, value := range rec {
		domainCities = append(domainCities, value.toDomain())
	}
	return domainCities, int(totalData), nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, id int) (cities.Domain, error) {
	rec := Cities{}
	err := nr.Conn.Where("id = ?", id).First(&rec).Error
	if err != nil {
		return cities.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, userDomain *cities.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
