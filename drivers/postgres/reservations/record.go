package reservations

import (
	"peduli-covid/businesses/reservations"
	"time"

	"gorm.io/gorm"
)

type Reservations struct {
	ID         int
	UserID     int
	HospitalID int
	BedtypeID  int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}

func (rec *Reservations) toDomain() reservations.Domain {
	return reservations.Domain{
		ID:         rec.ID,
		UserID:     rec.UserID,
		HospitalID: rec.HospitalID,
		BedtypeID:  rec.BedtypeID,
		Status:     rec.Status,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(reservationDomain reservations.Domain) *Reservations {
	return &Reservations{
		ID:         reservationDomain.ID,
		UserID:     reservationDomain.UserID,
		HospitalID: reservationDomain.HospitalID,
		BedtypeID:  reservationDomain.BedtypeID,
		Status:     reservationDomain.Status,
		CreatedAt:  reservationDomain.CreatedAt,
		UpdatedAt:  reservationDomain.UpdatedAt,
	}
}
