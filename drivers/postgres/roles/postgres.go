package roles

import (
	"context"
	"peduli-covid/businesses/roles"

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

func (nr *PostgresRepository) Fetch(ctx context.Context, page, perpage int) ([]roles.Domain, int, error) {
	rec := []Roles{}

	offset := (page - 1) * perpage
	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
	if err != nil {
		return []roles.Domain{}, 0, err
	}

	var totalData int64
	err = nr.Conn.Count(&totalData).Error
	if err != nil {
		return []roles.Domain{}, 0, err
	}

	var domainRoles []roles.Domain
	for _, value := range rec {
		domainRoles = append(domainRoles, value.toDomain())
	}
	return domainRoles, int(totalData), nil
}

func (nr *PostgresRepository) GetByID(RoleID int) (roles.Domain, error) {
	rec := Roles{}
	err := nr.Conn.Where("id = ?", RoleID).First(&rec).Error
	if err != nil {
		return roles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) GetByEmail(ctx context.Context, email string) (roles.Domain, error) {
	rec := Roles{}
	err := nr.Conn.Where("email = ?", email).First(&rec).Error
	if err != nil {
		return roles.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *PostgresRepository) Store(ctx context.Context, RoleDomain *roles.Domain) error {
	rec := fromDomain(*RoleDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
