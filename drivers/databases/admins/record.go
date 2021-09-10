package admins

import (
	"peduli-covid/businesses/admins"
	"time"
)

type Admins struct {
	ID         int
	RoleID     int
	HospitalID int
	Email      string
	Password   string
	Phone      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (rec *Admins) toDomain() admins.Domain {
	return admins.Domain{
		Id:         rec.ID,
		RoleID:     rec.RoleID,
		HospitalID: rec.HospitalID,
		Email:      rec.Email,
		Password:   rec.Password,
		Phone:      rec.Phone,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(userDomain admins.Domain) *Admins {
	return &Admins{
		ID:         userDomain.Id,
		RoleID:     userDomain.RoleID,
		HospitalID: userDomain.HospitalID,
		Email:      userDomain.Email,
		Password:   userDomain.Password,
		Phone:      userDomain.Phone,
		CreatedAt:  userDomain.CreatedAt,
		UpdatedAt:  userDomain.UpdatedAt,
	}
}
