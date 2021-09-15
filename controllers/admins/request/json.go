package request

import "peduli-covid/businesses/users"

type Admins struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	Phone      string `json:"phone"`
	RoleID     int    `json:"role_id"`
	HospitalID int    `json:"hospital_id"`
}

func (req *Admins) ToDomain() *users.Domain {
	return &users.Domain{
		Email:      req.Email,
		Password:   req.Password,
		Phone:      req.Phone,
		RoleID:     req.RoleID,
		HospitalID: req.HospitalID,
	}
}
