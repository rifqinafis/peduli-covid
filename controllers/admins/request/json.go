package request

import "peduli-covid/businesses/admins"

type Admins struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	RoleID     int    `json:"role_id"`
	HospitalID int    `json:"hospital_id"`
}

func (req *Admins) ToDomain() *admins.Domain {
	return &admins.Domain{
		Email:      req.Email,
		Password:   req.Password,
		Phone:      req.Phone,
		RoleID:     req.RoleID,
		HospitalID: req.HospitalID,
	}
}
