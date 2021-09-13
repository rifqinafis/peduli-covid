package admins

import (
	"net/http"
	"peduli-covid/businesses/admins"
	controller "peduli-covid/controllers"
	"peduli-covid/controllers/admins/request"

	echo "github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUsecase admins.Usecase
}

func NewAdminController(uc admins.Usecase) *AdminController {
	return &AdminController{
		adminUsecase: uc,
	}
}

func (ctrl *AdminController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.adminUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *AdminController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.adminUsecase.Login(ctx, req.Email, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, response)
}
