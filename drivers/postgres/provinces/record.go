package provinces

import (
	"peduli-covid/businesses/provinces"
	"time"

	"gorm.io/gorm"
)

type Provinces struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Provinces) toDomain() provinces.Domain {
	return provinces.Domain{
		ID:        rec.ID,
		Code:      rec.Code,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}

func fromDomain(provinceDomain provinces.Domain) *Provinces {
	return &Provinces{
		ID:        provinceDomain.ID,
		Code:      provinceDomain.Code,
		Name:      provinceDomain.Name,
		CreatedAt: provinceDomain.CreatedAt,
		UpdatedAt: provinceDomain.UpdatedAt,
		DeletedAt: provinceDomain.DeletedAt,
	}
}
