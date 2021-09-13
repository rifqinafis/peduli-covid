package response

import (
	"peduli-covid/businesses/admins"
	"time"
)

type Admins struct {
	ID         int       `json:"id"`
	RoleID     int       `json:"role_id"`
	HospitalID int       `json:"hospital_id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain admins.Domain) Admins {
	return Admins{
		ID:         domain.ID,
		RoleID:     domain.RoleID,
		HospitalID: domain.HospitalID,
		Email:      domain.Email,
		Password:   domain.Password,
		Phone:      domain.Phone,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
