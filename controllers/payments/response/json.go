package response

import (
	"peduli-covid/businesses/payments"
	"time"
)

type Payments struct {
	ID              int       `json:"id"`
	PaymentMethodID int       `json:"paymentmethod_id"`
	ReservationID   int       `json:"reservation_id"`
	Amount          float64   `json:"amount"`
	Date            string    `json:"date"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func FromDomain(domain payments.Domain) Payments {
	return Payments{
		ID:              domain.ID,
		PaymentMethodID: domain.PaymentMethodID,
		ReservationID:   domain.ReservationID,
		Amount:          domain.Amount,
		Date:            domain.Date,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}
