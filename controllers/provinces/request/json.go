package request

import "peduli-covid/businesses/provinces"

type Provinces struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func (req *Provinces) ToDomain() *provinces.Domain {
	return &provinces.Domain{
		Name: req.Name,
		Code: req.Code,
	}
}
