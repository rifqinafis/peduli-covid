package admins

import (
	"context"
	"time"
)

type Domain struct {
	ID         int       `json:"id"`
	RoleID     int       `json:"role_id"`
	HospitalID int       `json:"hospital_id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Usecase interface {
	Login(ctx context.Context, email, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
