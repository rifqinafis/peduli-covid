package invoices

import (
	"peduli-covid/businesses/invoices"
	"time"

	"gorm.io/gorm"
)

type Invoices struct {
	ID            int
	ReservationID int
	Code          string
	Total         float64
	AdminFee      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (rec *Invoices) toDomain() invoices.Domain {
	return invoices.Domain{
		ID:            rec.ID,
		ReservationID: rec.ReservationID,
		Code:          rec.Code,
		Total:         rec.Total,
		AdminFee:      rec.AdminFee,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
	}
}

func fromDomain(invoiceDomain invoices.Domain) *Invoices {
	return &Invoices{
		ID:            invoiceDomain.ID,
		ReservationID: invoiceDomain.ReservationID,
		Code:          invoiceDomain.Code,
		Total:         invoiceDomain.Total,
		AdminFee:      invoiceDomain.AdminFee,
		CreatedAt:     invoiceDomain.CreatedAt,
		UpdatedAt:     invoiceDomain.UpdatedAt,
	}
}
