package users

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/helpers/encrypt"
	"peduli-covid/helpers/messages"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Login(ctx context.Context, Email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(Email) == "" && strings.TrimSpace(password) == "" {
		return "", messages.ErrEmailPasswordNotFound
	}

	userDomain, err := uc.userRepository.GetByEmail(ctx, Email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, userDomain.Password) {
		return "", messages.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(userDomain.ID)
	return token, nil
}

func (uc *userUsecase) Store(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	userDomain.RoleID = 2

	existedUser, err := uc.userRepository.GetByEmail(ctx, userDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return messages.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)
	if err != nil {
		return messages.ErrInternalServer
	}

	err = uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return err
	}

	return nil
}
