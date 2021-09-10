package reservations

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/invoices"
	"time"
)

type reservationUsecase struct {
	reservationRepository Repository
	contextTimeout        time.Duration
	jwtAuth               *middleware.ConfigJWT
}

func NewReservationUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &reservationUsecase{
		reservationRepository: ur,
		jwtAuth:               jwtauth,
		contextTimeout:        timeout,
	}
}

func (uc *reservationUsecase) Store(ctx context.Context, reservationDomain *Domain) (res invoices.Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	data, err := uc.reservationRepository.Store(ctx, reservationDomain)
	if err != nil {
		return res, err
	}

	return res, nil
}
