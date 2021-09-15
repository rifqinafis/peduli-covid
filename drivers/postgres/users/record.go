package users

import (
	"peduli-covid/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID         int
	RoleID     int
	HospitalID int
	Email      string
	Password   string
	Phone      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		ID:         rec.ID,
		RoleID:     rec.RoleID,
		HospitalID: rec.HospitalID,
		Email:      rec.Email,
		Password:   rec.Password,
		Phone:      rec.Phone,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
		DeletedAt:  rec.DeletedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:         userDomain.ID,
		RoleID:     userDomain.RoleID,
		HospitalID: userDomain.HospitalID,
		Email:      userDomain.Email,
		Password:   userDomain.Password,
		Phone:      userDomain.Phone,
		CreatedAt:  userDomain.CreatedAt,
		UpdatedAt:  userDomain.UpdatedAt,
		DeletedAt:  userDomain.DeletedAt,
	}
}
