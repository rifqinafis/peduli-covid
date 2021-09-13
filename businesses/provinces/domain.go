package provinces

import (
	"context"
	"time"
)

type Domain struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Usecase interface {
	Store(ctx context.Context, data *Domain) error
	StoreFromAPI(ctx context.Context) error
}

type Repository interface {
	Store(ctx context.Context, data *Domain) error
	GetByCode(ctx context.Context, code string) (Domain, error)
}
