package request

import "peduli-covid/businesses/users"

type Users struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	RoleID   int    `json:"role_id"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		RoleID:   req.RoleID,
	}
}
