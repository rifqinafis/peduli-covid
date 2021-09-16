package notifications

import (
	"context"
	"time"
)

type notificationUsecase struct {
	notificationRepository Repository
	contextTimeout         time.Duration
}

func NewNotificationUsecase(ur Repository, timeout time.Duration) Usecase {
	return &notificationUsecase{
		notificationRepository: ur,
		contextTimeout:         timeout,
	}
}

func (uc *notificationUsecase) FindByUserID(ctx context.Context, userID int) ([]Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err := uc.notificationRepository.FindByUserID(ctx, userID)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (uc *notificationUsecase) Store(ctx context.Context, notificationDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.notificationRepository.Store(ctx, notificationDomain)
	if err != nil {
		return err
	}

	return nil
}
