package invoices_test

import (
	"context"
	"errors"
	"os"
	"peduli-covid/businesses/invoices"
	invoice_mock "peduli-covid/businesses/invoices/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	invoiceRepository invoice_mock.Repository
	invoiceUsecase    invoices.Usecase
)

func setup() {
	invoiceUsecase = invoices.NewInvoiceUsecase(&invoiceRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByUserID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		invoiceDomain := invoices.Domain{
			ID:            1,
			ReservationID: 2,
			Code:          "INVYTGS74TD",
			Total:         150000,
			AdminFee:      0,
		}
		invoiceRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]invoices.Domain{invoiceDomain}, nil).Once()

		result, err := invoiceUsecase.FindByUserID(context.Background(), 6)

		assert.Nil(t, err)
		assert.Equal(t, []invoices.Domain{invoiceDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		invoiceRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]invoices.Domain{}, errors.New("error")).Once()

		result, err := invoiceUsecase.FindByUserID(context.Background(), 6)

		assert.Equal(t, []invoices.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		invoiceDomain := invoices.Domain{
			ID:            1,
			ReservationID: 2,
			Code:          "INVYTGS74TD",
			Total:         150000,
			AdminFee:      0,
		}
		invoiceRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(invoiceDomain, nil).Once()

		result, err := invoiceUsecase.GetByID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, invoiceDomain, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		invoiceRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(invoices.Domain{}, errors.New("error")).Once()

		result, err := invoiceUsecase.GetByID(context.Background(), 1)

		assert.Equal(t, invoices.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		invoiceDomain := invoices.Domain{
			ID:            1,
			ReservationID: 2,
			Code:          "INVYTGS74TD",
			Total:         150000,
			AdminFee:      0,
		}
		invoiceRepository.On("Store", mock.Anything, mock.Anything).Return(invoiceDomain, nil).Once()

		result, err := invoiceUsecase.Store(context.Background(), &invoiceDomain)

		assert.Nil(t, err)
		assert.Equal(t, invoiceDomain, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		invoiceRepository.On("Store", mock.Anything, mock.Anything).Return(invoices.Domain{}, errors.New("error")).Once()

		result, err := invoiceUsecase.Store(context.Background(), &invoices.Domain{})

		assert.Equal(t, invoices.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}
