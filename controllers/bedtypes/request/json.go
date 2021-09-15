package request

import "peduli-covid/businesses/bedtypes"

type Bedtypes struct {
	HospitalID   int    `json:"hospital_id"`
	Name         string `json:"name"`
	BedAvailable int    `json:"bed_available"`
	BedEmpty     int    `json:"bed_empty"`
}

func (req *Bedtypes) ToDomain() *bedtypes.Domain {
	return &bedtypes.Domain{
		HospitalID:   req.HospitalID,
		Name:         req.Name,
		BedAvailable: req.BedAvailable,
		BedEmpty:     req.BedEmpty,
	}
}
