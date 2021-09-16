package notifications

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int            `json:"id"`
	UserID    int            `json:"user_id"`
	Code      string         `json:"code"`
	Details   string         `json:"details"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
}

type Repository interface {
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
}
