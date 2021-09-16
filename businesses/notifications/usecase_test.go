package notifications_test

import (
	"context"
	"errors"
	"os"
	"peduli-covid/businesses/notifications"
	notification_mock "peduli-covid/businesses/notifications/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	notificationRepository notification_mock.Repository
	notificationUsecase    notifications.Usecase
)

func setup() {
	notificationUsecase = notifications.NewNotificationUsecase(&notificationRepository, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestFindByUserID(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		notificationDomain := notifications.Domain{
			ID:      1,
			UserID:  4,
			Code:    "reservation",
			Details: "user test@gmail.com telah melakukan reservasi kamar Isolasi Tanpa Tekanan Negatif",
		}
		notificationRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]notifications.Domain{notificationDomain}, nil).Once()

		result, err := notificationUsecase.FindByUserID(context.Background(), 6)

		assert.Nil(t, err)
		assert.Equal(t, []notifications.Domain{notificationDomain}, result)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		notificationRepository.On("FindByUserID", mock.Anything, mock.AnythingOfType("int")).Return([]notifications.Domain{}, errors.New("error")).Once()

		result, err := notificationUsecase.FindByUserID(context.Background(), 6)

		assert.Equal(t, []notifications.Domain{}, result)
		assert.Equal(t, errors.New("error"), err)
	})
}

func TestStore(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {
		notificationDomain := notifications.Domain{
			ID:      1,
			UserID:  4,
			Code:    "reservation",
			Details: "user test@gmail.com telah melakukan reservasi kamar Isolasi Tanpa Tekanan Negatif",
		}
		notificationRepository.On("Store", mock.Anything, mock.Anything).Return(nil).Once()

		err := notificationUsecase.Store(context.Background(), &notificationDomain)

		assert.Equal(t, nil, err)
	})

	t.Run("test case 2, repository error test", func(t *testing.T) {
		notificationRepository.On("Store", mock.Anything, mock.Anything).Return(errors.New("error")).Once()

		err := notificationUsecase.Store(context.Background(), &notifications.Domain{})

		assert.Equal(t, errors.New("error"), err)
	})
}
