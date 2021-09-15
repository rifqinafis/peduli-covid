package cities

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           int            `json:"id"`
	ProvinceCode string         `json:"province_code"`
	Code         string         `json:"code"`
	Name         string         `json:"name"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	FindAll(ctx context.Context) ([]Domain, error)
	FindByProvinceCode(ctx context.Context, provinceCode string) ([]Domain, error)
	StoreFromAPI(ctx context.Context) error
}

type Repository interface {
	FindAll(ctx context.Context) ([]Domain, error)
	FindByProvinceCode(ctx context.Context, provinceCode string) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
