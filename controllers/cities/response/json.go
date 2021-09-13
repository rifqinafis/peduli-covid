package response

import (
	"peduli-covid/businesses/cities"
	"time"
)

type Cities struct {
	ID           int       `json:"id"`
	ProvinceCode string    `json:"province_code"`
	Name         string    `json:"name"`
	Code         string    `json:"code"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromDomain(domain cities.Domain) Cities {
	return Cities{
		ID:           domain.ID,
		ProvinceCode: domain.ProvinceCode,
		Name:         domain.Name,
		Code:         domain.Code,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
	}
}
