package request

import "peduli-covid/businesses/reservations"

type Reservations struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	HospitalID int    `json:"hospital_id"`
	BedtypeID  int    `json:"bedtype_id"`
	Status     string `json:"status"`
}

func (req *Reservations) ToDomain() *reservations.Domain {
	return &reservations.Domain{
		ID:         req.ID,
		UserID:     req.UserID,
		HospitalID: req.HospitalID,
		BedtypeID:  req.BedtypeID,
		Status:     req.Status,
	}
}
