package admins

import (
	"context"
	"peduli-covid/app/middleware"
	"peduli-covid/helpers/encrypt"
	"peduli-covid/helpers/messages"
	"strings"
	"time"
)

type adminUsecase struct {
	adminRepository Repository
	contextTimeout  time.Duration
	jwtAuth         *middleware.ConfigJWT
}

func NewAdminUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &adminUsecase{
		adminRepository: ur,
		jwtAuth:         jwtauth,
		contextTimeout:  timeout,
	}
}

func (uc *adminUsecase) Login(ctx context.Context, Email, password string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if strings.TrimSpace(Email) == "" && strings.TrimSpace(password) == "" {
		return "", messages.ErrEmailPasswordNotFound
	}

	adminDomain, err := uc.adminRepository.GetByEmail(ctx, Email)
	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, adminDomain.Password) {
		return "", messages.ErrInternalServer
	}

	token := uc.jwtAuth.GenerateToken(adminDomain.ID)
	return token, nil
}

func (uc *adminUsecase) Store(ctx context.Context, adminDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	adminDomain.RoleID = 1

	existedUser, err := uc.adminRepository.GetByEmail(ctx, adminDomain.Email)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return err
		}
	}
	if existedUser != (Domain{}) {
		return messages.ErrDuplicateData
	}

	adminDomain.Password, err = encrypt.Hash(adminDomain.Password)
	if err != nil {
		return messages.ErrInternalServer
	}

	err = uc.adminRepository.Store(ctx, adminDomain)
	if err != nil {
		return err
	}

	return nil
}
