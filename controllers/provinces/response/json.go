package response

import (
	"peduli-covid/businesses/provinces"
	"time"
)

type Provinces struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain provinces.Domain) Provinces {
	return Provinces{
		ID:        domain.ID,
		Name:      domain.Name,
		Code:      domain.Code,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
