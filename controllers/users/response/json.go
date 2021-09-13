package response

import (
	"peduli-covid/businesses/users"
	"time"
)

type Users struct {
	ID        int       `json:"id"`
	RoleID    int       `json:"role_id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain users.Domain) Users {
	return Users{
		ID:        domain.ID,
		RoleID:    domain.RoleID,
		Email:     domain.Email,
		Password:  domain.Password,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
