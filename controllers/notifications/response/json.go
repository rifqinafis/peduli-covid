package response

import (
	"peduli-covid/businesses/notifications"
	"time"
)

type Notifications struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Code      string    `json:"code"`
	Details   string    `json:"details"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain notifications.Domain) Notifications {
	return Notifications{
		ID:        domain.ID,
		UserID:    domain.UserID,
		Code:      domain.Code,
		Details:   domain.Details,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
