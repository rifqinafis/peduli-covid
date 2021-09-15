package provinces

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int            `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	FindAll(ctx context.Context) ([]Domain, error)
	StoreFromAPI(ctx context.Context) error
}

type Repository interface {
	FindAll(ctx context.Context) ([]Domain, error)
	GetByCode(ctx context.Context, code string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
