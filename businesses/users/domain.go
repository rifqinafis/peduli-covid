package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int       `json:"id"`
	RoleID    int       `json:"role_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
