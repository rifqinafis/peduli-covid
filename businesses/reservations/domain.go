package reservations

import (
	"context"
	"peduli-covid/businesses/invoices"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         int            `json:"id"`
	UserID     int            `json:"user_id"`
	HospitalID int            `json:"hospital_id"`
	BedtypeID  int            `json:"bedtype_id"`
	Status     string         `json:"status"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
	FindByAdminID(ctx context.Context, adminID int) ([]Domain, error)
	UpdateStatus(ctx context.Context, data *Domain) error
	UpdateStatusDone(ctx context.Context, data *Domain) error
	Store(ctx context.Context, data *Domain) (invoices.Domain, error)
}

type Repository interface {
	FindAll(ctx context.Context) ([]Domain, error)
	FindByUserID(ctx context.Context, userID int) ([]Domain, error)
	FindByHospitalID(ctx context.Context, hospitalID int) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	UpdateStatus(ctx context.Context, data *Domain) error
	Store(ctx context.Context, data *Domain) (Domain, error)
}
