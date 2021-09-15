package response

import (
	"peduli-covid/businesses/reservations"
	"time"
)

type Reservations struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	HospitalID int       `json:"hospital_id"`
	BedtypeID  int       `json:"bedtype_id"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(domain reservations.Domain) Reservations {
	return Reservations{
		ID:         domain.ID,
		UserID:     domain.UserID,
		HospitalID: domain.HospitalID,
		BedtypeID:  domain.BedtypeID,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}
