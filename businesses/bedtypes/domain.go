package bedtypes

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           int            `json:"id"`
	HospitalID   int            `json:"hospital_id"`
	Name         string         `json:"name"`
	BedAvailable int            `json:"bed_available"`
	BedEmpty     int            `json:"bed_empty"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	FindByHospitalID(ctx context.Context, hospitalID int) ([]Domain, error)
	StoreFromAPI(ctx context.Context) error
}

type Repository interface {
	FindByHospitalID(ctx context.Context, hospitalID int) ([]Domain, error)
	FindByTitleAndHospitalID(ctx context.Context, title string, id int) (Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	UpdateBedEmpty(ctx context.Context, data *Domain) error
	UpdateAvailableBed(ctx context.Context, data *Domain) error
	Store(ctx context.Context, data *Domain) error
}
