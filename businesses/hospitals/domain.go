package hospitals

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int            `json:"id"`
	CityID    int            `json:"city_id"`
	Name      string         `json:"name"`
	Address   string         `json:"address"`
	Phone     string         `json:"phone"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	FindAll(ctx context.Context) ([]Domain, error)
	FindByCityID(ctx context.Context, cityID int) ([]Domain, error)
	StoreFromAPI(ctx context.Context) error
}

type Repository interface {
	FindAll(ctx context.Context) ([]Domain, error)
	FindByCityID(ctx context.Context, cityID int) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
