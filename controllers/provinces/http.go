package provinces

import (
	"net/http"
	"peduli-covid/businesses/provinces"
	controller "peduli-covid/controllers"

	echo "github.com/labstack/echo/v4"
)

type ProvinceController struct {
	provinceUsecase provinces.Usecase
}

func NewProvinceController(uc provinces.Usecase) *ProvinceController {
	return &ProvinceController{
		provinceUsecase: uc,
	}
}

func (ctrl *ProvinceController) StoreFromAPI(c echo.Context) error {
	ctx := c.Request().Context()

	err := ctrl.provinceUsecase.StoreFromAPI(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *ProvinceController) FindAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.provinceUsecase.FindAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
