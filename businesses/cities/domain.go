package cities

import (
	"context"
	"time"
)

type Domain struct {
	ID           int       `json:"id"`
	ProvinceCode string    `json:"province_code"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	StoreFromAPI(ctx context.Context) error
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByID(ctx context.Context, id int) (Domain, error)
}
