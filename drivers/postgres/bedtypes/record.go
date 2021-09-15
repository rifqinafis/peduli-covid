package bedtypes

import (
	"peduli-covid/businesses/bedtypes"
	"time"

	"gorm.io/gorm"
)

type Bedtypes struct {
	ID           int
	HospitalID   int
	Name         string
	BedAvailable int
	BedEmpty     int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (rec *Bedtypes) toDomain() bedtypes.Domain {
	return bedtypes.Domain{
		ID:           rec.ID,
		HospitalID:   rec.HospitalID,
		Name:         rec.Name,
		BedAvailable: rec.BedAvailable,
		BedEmpty:     rec.BedEmpty,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}

func fromDomain(bedtypeDomain bedtypes.Domain) *Bedtypes {
	return &Bedtypes{
		ID:           bedtypeDomain.ID,
		HospitalID:   bedtypeDomain.HospitalID,
		Name:         bedtypeDomain.Name,
		BedAvailable: bedtypeDomain.BedAvailable,
		BedEmpty:     bedtypeDomain.BedEmpty,
		CreatedAt:    bedtypeDomain.CreatedAt,
		UpdatedAt:    bedtypeDomain.UpdatedAt,
		DeletedAt:    bedtypeDomain.DeletedAt,
	}
}
