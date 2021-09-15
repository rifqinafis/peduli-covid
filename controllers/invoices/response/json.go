package response

import (
	"peduli-covid/businesses/invoices"
	"time"
)

type Invoices struct {
	ID            int       `json:"id"`
	ReservationID int       `json:"reservation_id"`
	Code          string    `json:"code"`
	Total         float64   `json:"total"`
	AdminFee      float64   `json:"admin_fee"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func FromDomain(domain invoices.Domain) Invoices {
	return Invoices{
		ID:            domain.ID,
		ReservationID: domain.ReservationID,
		Code:          domain.Code,
		Total:         domain.Total,
		AdminFee:      domain.AdminFee,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
