package payments

import (
	"context"
	"errors"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/reservations"
	consts "peduli-covid/helpers/const"
	"time"
)

type paymentUsecase struct {
	paymentRepository     Repository
	reservationRepository reservations.Repository
	contextTimeout        time.Duration
	jwtAuth               *middleware.ConfigJWT
}

func NewPaymentUsecase(ur Repository, reservationRepo reservations.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &paymentUsecase{
		paymentRepository:     ur,
		reservationRepository: reservationRepo,
		jwtAuth:               jwtauth,
		contextTimeout:        timeout,
	}
}

func (uc *paymentUsecase) FindByUserID(ctx context.Context, userID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.paymentRepository.FindByUserID(ctx, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *paymentUsecase) Store(ctx context.Context, paymentDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	paymentDomain.Date = time.Now().Format("2006-01-02")

	if paymentDomain.Amount != consts.RESERVATION_FEE {
		return errors.New("invalid amount")
	}

	err := uc.paymentRepository.Store(ctx, paymentDomain)
	if err != nil {
		return err
	}

	dataReservation := reservations.Domain{
		ID:     paymentDomain.ReservationID,
		Status: consts.STATUS_PAID,
	}

	err = uc.reservationRepository.UpdateStatus(ctx, &dataReservation)
	if err != nil {
		return err
	}

	return nil
}
