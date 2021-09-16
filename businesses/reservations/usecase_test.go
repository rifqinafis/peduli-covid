package reservations_test

import (
	"context"
	"errors"
	"os"
	bedtype_mock "peduli-covid/businesses/bedtypes/mocks"
	hospital_mock "peduli-covid/businesses/hospitals/mocks"
	invoice_mock "peduli-covid/businesses/invoices/mocks"
	notification_mock "peduli-covid/businesses/notifications/mocks"
	"peduli-covid/businesses/reservations"
	reservation_mock "peduli-covid/businesses/reservations/mocks"
	"peduli-covid/businesses/users"
	user_mock "peduli-covid/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	reservationRepository  reservation_mock.Repository
	invoiceRepository      invoice_mock.Repository
	userRepository         user_mock.Repository
	hospitalRepository     hospital_mock.Repository
	bedtypeRepository      bedtype_mock.Repository
	notificationRepository notification_mock.Repository
	reservationUsecase     reservations.Usecase
)

func setup() {
	reservationUsecase = reservations.NewReservationUsecase(&reservationRepository, &invoiceRepository, &userRepository, &hospitalRepository, &bedtypeRepository, &notificationRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByUserID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		reservationDomain := reservations.Domain{
			ID:         1,
			UserID:     1,
			HospitalID: 1,
			BedtypeID:  1,
			Status:     "paid",
		}
		reservationRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]reservations.Domain{reservationDomain}, nil).Once()

		result, err := reservationUsecase.FindByUserID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, []reservations.Domain{reservationDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		reservationRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]reservations.Domain{}, errors.New("error")).Once()

		result, err := reservationUsecase.FindByUserID(context.Background(), 1)

		assert.Equal(t, []reservations.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestFindByAdminID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		userDomain := users.Domain{
			ID:         1,
			HospitalID: 1,
			RoleID:     1,
			Email:      "test@gmail.com",
			Password:   "test123",
			Phone:      "081249848950",
		}
		userRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(userDomain, nil).Once()
		reservationDomain := reservations.Domain{
			ID:         1,
			UserID:     1,
			HospitalID: 1,
			BedtypeID:  1,
			Status:     "paid",
		}
		reservationRepository.On("FindByHospitalID", mock.Anything, mock.AnythingOfType("int")).Return([]reservations.Domain{reservationDomain}, nil).Once()

		result, err := reservationUsecase.FindByAdminID(context.Background(), 1)

		assert.Nil(t, err)
		assert.Equal(t, []reservations.Domain{reservationDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		userRepository.On("GetByID", mock.Anything, mock.AnythingOfType("int")).Return(users.Domain{}, errors.New("error")).Once()
		reservationRepository.On("FindByHospitalID", mock.Anything, mock.AnythingOfType("int")).Return([]reservations.Domain(nil), errors.New("error")).Once()

		result, err := reservationUsecase.FindByAdminID(context.Background(), 1)

		assert.Equal(t, []reservations.Domain(nil), result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestUpdateStatus(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		reservationDomain := reservations.Domain{
			ID:         1,
			UserID:     1,
			HospitalID: 1,
			BedtypeID:  1,
			Status:     "paid",
		}
		reservationRepository.On("UpdateStatus", mock.Anything, mock.Anything).Return(nil).Once()

		err := reservationUsecase.UpdateStatus(context.Background(), &reservationDomain)

		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		reservationRepository.On("UpdateStatus", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		err := reservationUsecase.UpdateStatus(context.Background(), &reservations.Domain{})

		assert.Equal(t, errors.New("error"), err)
	})
}
