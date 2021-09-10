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

func (uc *invoiceUsecase) Store(ctx context.Context, invoiceDomain *Domain) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.invoiceRepository.Store(ctx, invoiceDomain)
	if err != nil {
		return 0, err
	}

	return res.Id, nil
}
