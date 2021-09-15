package invoices

import (
	"context"
	"peduli-covid/app/middleware"
	"time"
)

type invoiceUsecase struct {
	invoiceRepository Repository
	contextTimeout    time.Duration
	jwtAuth           *middleware.ConfigJWT
}

func NewInvoiceUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &invoiceUsecase{
		invoiceRepository: ur,
		jwtAuth:           jwtauth,
		contextTimeout:    timeout,
	}
}

func (uc *invoiceUsecase) FindByUserID(ctx context.Context, userID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.invoiceRepository.FindByUserID(ctx, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *invoiceUsecase) GetByID(ctx context.Context, id int) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.invoiceRepository.GetByID(ctx, id)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *invoiceUsecase) Store(ctx context.Context, invoiceDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.invoiceRepository.Store(ctx, invoiceDomain)
	if err != nil {
		return res, err
	}

	return res, nil
}
