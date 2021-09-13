package request

import "peduli-covid/businesses/roles"

type Roles struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func (req *Roles) ToDomain() *roles.Domain {
	return &roles.Domain{
		Code: req.Code,
		Name: req.Name,
	}
}
