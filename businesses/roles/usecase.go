package roles

import (
	"peduli-covid/app/middleware"
	"time"
)

type roleUsecase struct {
	roleRepository Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewRoleUsecase(ur Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &roleUsecase{
		roleRepository: ur,
		jwtAuth:        jwtauth,
		contextTimeout: timeout,
	}
}

func (uc *roleUsecase) GetByID(id int) string {

	res, err := uc.roleRepository.GetByID(id)
	if err != nil {
		return res.Code
	}

	return res.Code
}
