package roles

import (
	"peduli-covid/businesses/roles"
)

type RoleController struct {
	roleUsecase roles.Usecase
}

func NewRoleController(uc roles.Usecase) *RoleController {
	return &RoleController{
		roleUsecase: uc,
	}
}

func (ctrl *RoleController) FindRole(id int) string {
	var role string
	data := ctrl.roleUsecase.GetByID(id)
	if data != "" {
		role = data
	}

	return role
}
