package request

import "peduli-covid/businesses/cities"

type Cities struct {
	ID           int    `json:"id"`
	ProvinceCode string `json:"province_code"`
	Name         string `json:"name"`
	Code         string `json:"code"`
}

func (req *Cities) ToDomain() *cities.Domain {
	return &cities.Domain{
		ID:           req.ID,
		ProvinceCode: req.ProvinceCode,
		Name:         req.Name,
		Code:         req.Code,
	}
}
