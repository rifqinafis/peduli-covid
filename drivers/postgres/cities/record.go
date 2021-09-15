package cities

import (
	"peduli-covid/businesses/cities"
	"time"

	"gorm.io/gorm"
)

type Cities struct {
	ID           int
	ProvinceCode string
	Code         string
	Name         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (rec *Cities) toDomain() cities.Domain {
	return cities.Domain{
		ID:           rec.ID,
		ProvinceCode: rec.ProvinceCode,
		Code:         rec.Code,
		Name:         rec.Name,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}

func fromDomain(cityDomain cities.Domain) *Cities {
	return &Cities{
		ID:           cityDomain.ID,
		ProvinceCode: cityDomain.ProvinceCode,
		Code:         cityDomain.Code,
		Name:         cityDomain.Name,
		CreatedAt:    cityDomain.CreatedAt,
		UpdatedAt:    cityDomain.UpdatedAt,
		DeletedAt:    cityDomain.DeletedAt,
	}
}
