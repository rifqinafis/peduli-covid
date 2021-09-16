package payments

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID              int            `json:"id"`
	PaymentMethodID int            `json:"paymentmethod_id"`
	ReservationID   int            `json:"reservation_id"`
	Amount          float64        `json:"amount"`
	Date            string         `json:"date"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	Store(ctx context.Context, userID int, data *Domain) error
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
}

type Repository interface {
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
}
