package payments

import (
	"context"
	"errors"
	"peduli-covid/businesses/notifications"
	"peduli-covid/businesses/reservations"
	"peduli-covid/businesses/users"
	consts "peduli-covid/helpers/const"
	"time"
)

type paymentUsecase struct {
	paymentRepository     Repository
	reservationRepository reservations.Repository
	userRepository        users.Repository
	notifRepository       notifications.Repository
	contextTimeout        time.Duration
}

func NewPaymentUsecase(ur Repository, reservationRepo reservations.Repository, userRepo users.Repository, notifRepo notifications.Repository, timeout time.Duration) Usecase {
	return &paymentUsecase{
		paymentRepository:     ur,
		reservationRepository: reservationRepo,
		userRepository:        userRepo,
		notifRepository:       notifRepo,
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

func (uc *paymentUsecase) Store(ctx context.Context, userID int, paymentDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	paymentDomain.Date = time.Now().Format("2006-01-02")

	if paymentDomain.Amount != consts.RESERVATION_FEE+consts.ADMIN_FEE {
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

	reservationData, err := uc.reservationRepository.GetByID(ctx, paymentDomain.ReservationID)
	if err != nil {
		return err
	}

	dataReserver, err := uc.userRepository.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	dataUser, err := uc.userRepository.FindByHospitalID(ctx, reservationData.HospitalID)
	if err != nil {
		return err
	}

	for _, data := range dataUser {
		notifReq := notifications.Domain{
			UserID:  data.ID,
			Code:    consts.NOTIF_PAYMENT,
			Details: "user " + dataReserver.Email + " telah melakukan pembayaran",
		}
		err = uc.notifRepository.Store(ctx, &notifReq)
		if err != nil {
			return err
		}
	}

	return nil
}
