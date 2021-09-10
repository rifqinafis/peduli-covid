package reservations

import (
	"context"
	"peduli-covid/businesses/invoices"
	"time"
)

type Domain struct {
	Id         int       `json:"id"`
	UserID     int       `json:"user_id"`
	HospitalID int       `json:"hospital_id"`
	BedType    string    `json:"bed_type"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) (invoices.Domain, error)
}

type Repository interface {
	Store(ctx context.Context, data *Domain) (Domain, error)
}
