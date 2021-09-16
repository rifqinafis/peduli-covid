package request

import "peduli-covid/businesses/notifications"

type Notifications struct {
	UserID  int    `json:"user_id"`
	Code    string `json:"code"`
	Details string `json:"details"`
}

func (req *Notifications) ToDomain() *notifications.Domain {
	return &notifications.Domain{
		UserID:  req.UserID,
		Code:    req.Code,
		Details: req.Details,
	}
}
