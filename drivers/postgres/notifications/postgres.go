package notifications

import (
	"context"
	"peduli-covid/businesses/notifications"

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

func (nr *PostgresRepository) FindByUserID(ctx context.Context, userID int) ([]notifications.Domain, error) {
	rec := []Notifications{}

	err := nr.Conn.Where("user_id = ?", userID).Find(&rec).Error
	if err != nil {
		return []notifications.Domain{}, err
	}

	var domainNotifications []notifications.Domain
	for _, value := range rec {
		domainNotifications = append(domainNotifications, value.toDomain())
	}
	return domainNotifications, nil
}

func (nr *PostgresRepository) Store(ctx context.Context, notificationDomain *notifications.Domain) error {
	rec := fromDomain(*notificationDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
