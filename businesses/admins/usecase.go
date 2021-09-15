package admins

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/users"
	"peduli-covid/helpers/encrypt"
	"peduli-covid/helpers/messages"
	"strings"
	"time"
)

type adminUsecase struct {
	userRepository users.Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewAdminUsecase(ur users.Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &adminUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *adminUsecase) Login(ctx context.Context, Email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(Email) == "" && strings.TrimSpace(password) == "" {
		return "", messages.ErrEmailPasswordNotFound
	}

	adminDomain, err := uc.userRepository.GetByEmail(ctx, Email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, adminDomain.Password) {
		return "", messages.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(adminDomain.ID)
	return token, nil
}

func (uc *adminUsecase) Store(ctx context.Context, adminDomain *users.Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	adminDomain.RoleID = 1

	existedUser, err := uc.userRepository.GetByEmail(ctx, adminDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (users.Domain{}) {
		return messages.ErrDuplicateData
	}

	adminDomain.Password, err = encrypt.Hash(adminDomain.Password)
	if err != nil {
		return messages.ErrInternalServer
	}

	err = uc.userRepository.Store(ctx, adminDomain)
	if err != nil {
		return err
	}

	return nil
}
