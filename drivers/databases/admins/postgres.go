package admins

import (
	"context"
	"peduli-covid/businesses/admins"

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

func (nr *PostgresRepository) Fetch(ctx context.Context, page, perpage int) ([]admins.Domain, int, error) {
	rec := []Admins{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []admins.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []admins.Domain{}, 0, err
	}

	var domainNews []admins.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, adminID int) (admins.Domain, error) {
	rec := Admins{}
	err := nr.Conn.Where("id = ?", adminID).First(&rec).Error
	if err != nil {
		return admins.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) GetByEmail(ctx context.Context, email string) (admins.Domain, error) {
	rec := Admins{}
	err := nr.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return admins.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, adminDomain *admins.Domain) error {
	rec := fromDomain(*adminDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
