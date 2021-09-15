package response

import (
	"peduli-covid/businesses/hospitals"
	"time"
)

type Hospitals struct {
	ID        int       `json:"id"`
	CityID    int       `json:"city_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain hospitals.Domain) Hospitals {
	return Hospitals{
		ID:        domain.ID,
		CityID:    domain.CityID,
		Name:      domain.Name,
		Address:   domain.Address,
		Phone:     domain.Phone,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
