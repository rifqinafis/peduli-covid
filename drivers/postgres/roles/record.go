package roles

import (
	"peduli-covid/businesses/roles"
	"time"
)

type Roles struct {
	ID        int
	Name      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Roles) toDomain() roles.Domain {
	return roles.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		Code:      rec.Code,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(roleDomain roles.Domain) *Roles {
	return &Roles{
		ID:        roleDomain.ID,
		Name:      roleDomain.Name,
		Code:      roleDomain.Code,
		CreatedAt: roleDomain.CreatedAt,
		UpdatedAt: roleDomain.UpdatedAt,
	}
}
