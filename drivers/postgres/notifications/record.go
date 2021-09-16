package notifications

import (
	"peduli-covid/businesses/notifications"
	"time"

	"gorm.io/gorm"
)

type Notifications struct {
	ID        int
	UserID    int
	Code      string
	Details   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Notifications) toDomain() notifications.Domain {
	return notifications.Domain{
		ID:        rec.ID,
		UserID:    rec.UserID,
		Code:      rec.Code,
		Details:   rec.Details,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(notificationDomain notifications.Domain) *Notifications {
	return &Notifications{
		ID:        notificationDomain.ID,
		UserID:    notificationDomain.UserID,
		Code:      notificationDomain.Code,
		Details:   notificationDomain.Details,
		CreatedAt: notificationDomain.CreatedAt,
		UpdatedAt: notificationDomain.UpdatedAt,
		DeletedAt: notificationDomain.DeletedAt,
	}
}
