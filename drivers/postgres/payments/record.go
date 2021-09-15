package payments

import (
	"peduli-covid/businesses/payments"
	"time"

	"gorm.io/gorm"
)

type Payments struct {
	ID              int
	PaymentmethodID int
	ReservationID   int
	Amount          float64
	Date            string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}

func (rec *Payments) toDomain() payments.Domain {
	return payments.Domain{
		ID:              rec.ID,
		PaymentMethodID: rec.PaymentmethodID,
		ReservationID:   rec.ReservationID,
		Amount:          rec.Amount,
		Date:            rec.Date,
		CreatedAt:       rec.CreatedAt,
		UpdatedAt:       rec.UpdatedAt,
	}
}

func fromDomain(paymentDomain payments.Domain) *Payments {
	return &Payments{
		ID:              paymentDomain.ID,
		PaymentmethodID: paymentDomain.PaymentMethodID,
		ReservationID:   paymentDomain.ReservationID,
		Amount:          paymentDomain.Amount,
		Date:            paymentDomain.Date,
		CreatedAt:       paymentDomain.CreatedAt,
		UpdatedAt:       paymentDomain.UpdatedAt,
	}
}
