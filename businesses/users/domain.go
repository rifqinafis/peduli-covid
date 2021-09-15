package users

import (
	"context"
	"peduli-covid/businesses/roles"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID         int            `json:"id"`
	HospitalID int            `json:"hospital_id"`
	RoleID     int            `json:"role_id"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	Phone      string         `json:"phone"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

type Usecase interface {
	GetByID(ctx context.Context, id int) (roles.Domain, error)
	Login(ctx context.Context, email, password string) (string, error)
	Store(ctx context.Context, data *Domain) error
}

type Repository interface {
	GetByID(ctx context.Context, id int) (Domain, error)
	GetByEmail(ctx context.Context, email string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
