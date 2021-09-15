package hospitals

import (
	"peduli-covid/businesses/hospitals"
	"time"

	"gorm.io/gorm"
)

type Hospitals struct {
	ID        int
	CityID    int
	Name      string
	Address   string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (rec *Hospitals) toDomain() hospitals.Domain {
	return hospitals.Domain{
		ID:        rec.ID,
		CityID:    rec.CityID,
		Name:      rec.Name,
		Address:   rec.Address,
		Phone:     rec.Phone,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(hospitalDomain hospitals.Domain) *Hospitals {
	return &Hospitals{
		ID:        hospitalDomain.ID,
		CityID:    hospitalDomain.CityID,
		Name:      hospitalDomain.Name,
		Address:   hospitalDomain.Address,
		Phone:     hospitalDomain.Phone,
		CreatedAt: hospitalDomain.CreatedAt,
		UpdatedAt: hospitalDomain.UpdatedAt,
	}
}
