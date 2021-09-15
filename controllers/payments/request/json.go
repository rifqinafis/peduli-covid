package request

import "peduli-covid/businesses/payments"

type Payments struct {
	PaymentMethodID int     `json:"paymentmethod_id"`
	ReservationID   int     `json:"reservation_id"`
	Amount          float64 `json:"amount"`
}

func (req *Payments) ToDomain() *payments.Domain {
	return &payments.Domain{
		PaymentMethodID: req.PaymentMethodID,
		ReservationID:   req.ReservationID,
		Amount:          req.Amount,
	}
}
