package invoices

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            int            `json:"id"`
	ReservationID int            `json:"reservation_id"`
	Code          string         `json:"code"`
	Total         float64        `json:"total"`
	AdminFee      float64        `json:"admin_fee"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) (Domain, error)
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
	Store(ctx context.Context, data *Domain) (Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}
