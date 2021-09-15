package request

import "peduli-covid/businesses/hospitals"

type Hospitals struct {
	ID      int    `json:"id"`
	CityID  int    `json:"city_id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func (req *Hospitals) ToDomain() *hospitals.Domain {
	return &hospitals.Domain{
		ID:      req.ID,
		CityID:  req.CityID,
		Name:    req.Name,
		Address: req.Address,
		Phone:   req.Phone,
	}
}
