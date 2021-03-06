package reservations

import (
	"context"
	"errors"
	"peduli-covid/businesses/bedtypes"
	"peduli-covid/businesses/hospitals"
	"peduli-covid/businesses/invoices"
	"peduli-covid/businesses/notifications"
	"peduli-covid/businesses/users"
	consts "peduli-covid/helpers/const"
	"peduli-covid/helpers/str"
	"time"
)

type reservationUsecase struct {
	reservationRepository  Repository
	invoiceRepository      invoices.Repository
	userRepository         users.Repository
	hospitalRepository     hospitals.Repository
	bedtypeRepository      bedtypes.Repository
	notificationRepository notifications.Repository
	contextTimeout         time.Duration
}

func NewReservationUsecase(ur Repository, invoiceRepo invoices.Repository, userRepo users.Repository, hospitalRepo hospitals.Repository, bedtypeRepo bedtypes.Repository, notifRepo notifications.Repository, timeout time.Duration) Usecase {
	return &reservationUsecase{
		reservationRepository:  ur,
		invoiceRepository:      invoiceRepo,
		userRepository:         userRepo,
		hospitalRepository:     hospitalRepo,
		bedtypeRepository:      bedtypeRepo,
		notificationRepository: notifRepo,
		contextTimeout:         timeout,
	}
}

func (uc *reservationUsecase) FindByUserID(ctx context.Context, userID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.reservationRepository.FindByUserID(ctx, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *reservationUsecase) FindByAdminID(ctx context.Context, userID int) (res []Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	dataAdmin, err := uc.userRepository.GetByID(ctx, userID)
	if err != nil {
		return res, err
	}

	res, err = uc.reservationRepository.FindByHospitalID(ctx, dataAdmin.HospitalID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *reservationUsecase) UpdateStatus(ctx context.Context, reservationDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.reservationRepository.UpdateStatus(ctx, reservationDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *reservationUsecase) UpdateStatusDone(ctx context.Context, reservationDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	dataReservation, err := uc.reservationRepository.GetByID(ctx, reservationDomain.ID)
	if err != nil {
		return err
	}

	if dataReservation.Status != consts.STATUS_PAID {
		return errors.New("status unpaid")
	}

	reservationDomain.Status = consts.STATUS_DONE

	err = uc.reservationRepository.UpdateStatus(ctx, reservationDomain)
	if err != nil {
		return err
	}

	dataBedtype, err := uc.bedtypeRepository.GetByID(ctx, dataReservation.BedtypeID)
	if err != nil {
		return err
	}

	domainBedtype := bedtypes.Domain{
		ID:       dataBedtype.ID,
		BedEmpty: dataBedtype.BedEmpty - 1,
	}

	err = uc.bedtypeRepository.UpdateBedEmpty(ctx, &domainBedtype)
	if err != nil {
		return err
	}

	dataUser, err := uc.userRepository.GetByID(ctx, dataReservation.UserID)
	if err != nil {
		return err
	}

	notifReq := notifications.Domain{
		UserID:  dataUser.ID,
		Code:    consts.NOTIF_VERIFIED,
		Details: "admin telah memverifikasi reservasi anda dengan kamar " + dataBedtype.Name,
	}
	err = uc.notificationRepository.Store(ctx, &notifReq)
	if err != nil {
		return err
	}

	return nil
}

func (uc *reservationUsecase) Store(ctx context.Context, reservationDomain *Domain) (res invoices.Domain, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	dataBedtype, err := uc.bedtypeRepository.GetByID(ctx, reservationDomain.BedtypeID)
	if err != nil {
		return res, err
	}

	if dataBedtype.BedAvailable == 0 {
		return res, errors.New("no available bed")
	}

	reservationDomain.Status = consts.STATUS_PENDING

	data, err := uc.reservationRepository.Store(ctx, reservationDomain)
	if err != nil {
		return res, err
	}

	bedtypeDomain := bedtypes.Domain{
		ID:           dataBedtype.ID,
		BedAvailable: dataBedtype.BedAvailable - 1,
	}

	err = uc.bedtypeRepository.UpdateAvailableBed(ctx, &bedtypeDomain)
	if err != nil {
		return res, err
	}

	req := invoices.Domain{
		ReservationID: data.ID,
		Code:          "INV" + str.RandAlphanumericString(8),
		Total:         consts.RESERVATION_FEE + consts.AMOUNT_CODE,
		AdminFee:      consts.ADMIN_FEE,
	}

	res, err = uc.invoiceRepository.Store(ctx, &req)
	if err != nil {
		return res, err
	}

	dataReserver, err := uc.userRepository.GetByID(ctx, reservationDomain.UserID)
	if err != nil {
		return res, err
	}

	dataUser, err := uc.userRepository.FindByHospitalID(ctx, reservationDomain.HospitalID)
	if err != nil {
		return res, err
	}

	for _, data := range dataUser {
		notifReq := notifications.Domain{
			UserID:  data.ID,
			Code:    consts.NOTIF_RESERVATION,
			Details: "user " + dataReserver.Email + " telah melakukan reservasi kamar " + dataBedtype.Name,
		}
		err = uc.notificationRepository.Store(ctx, &notifReq)
		if err != nil {
			return res, err
		}
	}

	return res, err
}
