package provinces

import (
	"peduli-covid/businesses/provinces"
	"time"
)

type Provinces struct {
	ID        int
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Provinces) toDomain() provinces.Domain {
	return provinces.Domain{
		ID:        rec.ID,
		Code:      rec.Code,
		Name:      rec.Name,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(provinceDomain provinces.Domain) *Provinces {
	return &Provinces{
		ID:        provinceDomain.ID,
		Code:      provinceDomain.Code,
		Name:      provinceDomain.Name,
		CreatedAt: provinceDomain.CreatedAt,
		UpdatedAt: provinceDomain.UpdatedAt,
	}
}
