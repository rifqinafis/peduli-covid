package roles

import (
	"time"
)

type roleUsecase struct {
	roleRepository Repository
	contextTimeout time.Duration
}

func NewRoleUsecase(ur Repository, timeout time.Duration) Usecase {
	return &roleUsecase{
		roleRepository: ur,
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
