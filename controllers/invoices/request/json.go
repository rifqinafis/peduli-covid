package request

import "peduli-covid/businesses/invoices"

type Invoices struct {
	ReservationID int     `json:"reservation_id"`
	Code          string  `json:"code"`
	Total         float64 `json:"total"`
	AdminFee      float64 `json:"admin_fee"`
}

func (req *Invoices) ToDomain() *invoices.Domain {
	return &invoices.Domain{
		ReservationID: req.ReservationID,
		Code:          req.Code,
		Total:         req.Total,
		AdminFee:      req.AdminFee,
	}
}
