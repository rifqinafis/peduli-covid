package users

import (
	"peduli-covid/businesses/users"
	"time"
)

type Users struct {
	ID        int
	RoleID    int
	Email     string
	Password  string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.ID,
		RoleID:    rec.RoleID,
		Email:     rec.Email,
		Password:  rec.Password,
		Phone:     rec.Phone,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.Id,
		RoleID:    userDomain.RoleID,
		Email:     userDomain.Email,
		Password:  userDomain.Password,
		Phone:     userDomain.Phone,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
