package response

import (
	"peduli-covid/businesses/bedtypes"
	"time"
)

type Bedtypes struct {
	ID           int       `json:"id"`
	HospitalID   int       `json:"hospital_id"`
	Name         string    `json:"name"`
	BedAvailable int       `json:"bed_available"`
	BedEmpty     int       `json:"bed_empty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain bedtypes.Domain) Bedtypes {
	return Bedtypes{
		ID:           domain.ID,
		HospitalID:   domain.HospitalID,
		Name:         domain.Name,
		BedAvailable: domain.BedAvailable,
		BedEmpty:     domain.BedEmpty,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
