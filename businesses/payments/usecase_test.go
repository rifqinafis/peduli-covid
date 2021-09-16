package payments_test

import (
	"context"
	"errors"
	"os"
	notification_mock "peduli-covid/businesses/notifications/mocks"
	"peduli-covid/businesses/payments"
	payment_mock "peduli-covid/businesses/payments/mocks"
	"peduli-covid/businesses/reservations"
	reservation_mock "peduli-covid/businesses/reservations/mocks"
	"peduli-covid/businesses/users"
	user_mock "peduli-covid/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	paymentRepository payment_mock.Repository
	reservationRepo   reservation_mock.Repository
	userRepo          user_mock.Repository
	notificationRepo  notification_mock.Repository
	paymentUsecase    payments.Usecase
)

func setup() {
	paymentUsecase = payments.NewPaymentUsecase(&paymentRepository, &reservationRepo, &userRepo, &notificationRepo, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByUserID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		paymentDomain := payments.Domain{
			ID:              1,
			PaymentMethodID: 1,
			ReservationID:   1,
			Amount:          150000,
			Date:            "2021-09-15",
		}
		paymentRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]payments.Domain{paymentDomain}, nil).Once()

		result, err := paymentUsecase.FindByUserID(context.Background(), 6)

		assert.Nil(t, err)
		assert.Equal(t, []payments.Domain{paymentDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		paymentRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]payments.Domain{}, errors.New("error")).Once()

		result, err := paymentUsecase.FindByUserID(context.Background(), 6)

		assert.Equal(t, []payments.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		paymentDomain := payments.Domain{
			ID:              1,
			PaymentMethodID: 1,
			ReservationID:   1,
			Amount:          151000,
			Date:            "2021-09-15",
		}
		paymentRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()
		reservationDomain := reservations.Domain{
			ID:         1,
			UserID:     1,
			HospitalID: 1,
			BedtypeID:  1,
			Status:     "paid",
		}
		reservationRepo.On("UpdateStatus", mock.Anything, mock.Anything).Return(nil).Once()
		reservationRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reservationDomain, nil).Once()
		userDomain := users.Domain{
			ID:         1,
			HospitalID: 1,
			RoleID:     1,
			Email:      "test@gmail.com",
			Password:   "test123",
			Phone:      "081249848950",
		}
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		userRepo.On("FindByHospitalID", mock.Anything, mock.AnythingOfType("int")).Return([]users.Domain{userDomain}, nil).Once()
		notificationRepo.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

		err := paymentUsecase.Store(context.Background(), 1, &paymentDomain)

		assert.Nil(t, err)
		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		paymentRepository.On("Store", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		reservationRepo.On("UpdateStatus", mock.Anything, mock.Anything).Return(errors.New("error")).Once()
		reservationRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(reservations.Domain{}, errors.New("error")).Once()
		userRepo.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("error")).Once()
		userRepo.On("FindByHospitalID", mock.Anything, mock.AnythingOfType("int")).Return([]users.Domain{}, errors.New("error")).Once()
		notificationRepo.On("Store", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		err := paymentUsecase.Store(context.Background(), 1, &payments.Domain{})

		assert.Equal(t, errors.New("invalid amount"), err)
	})
}
