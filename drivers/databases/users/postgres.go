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

func (nr *PostgresRepository) Fetch(ctx context.Context, page, perpage int) ([]users.Domain, int, error) {
	rec := []Users{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []users.Domain{}, 0, err
	}

	var domainNews []users.Domain
	for _, value := range rec {
		domainNews = append(domainNews, value.toDomain())
	}
	return domainNews, int(totalData), nil
}

func (nr *PostgresRepository) GetByID(ctx context.Context, userId int) (users.Domain, error) {
	rec := Users{}
	err := nr.Conn.Where("id = ?", userId).First(&rec).Error
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
