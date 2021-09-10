package invoices

import (
	"context"
	"time"
)

type Domain struct {
	Id            int       `json:"id"`
	ReservationID int       `json:"reservation_id"`
	Code          string    `json:"code"`
	Total         float64   `json:"total"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) (int, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (Domain, error)
}
