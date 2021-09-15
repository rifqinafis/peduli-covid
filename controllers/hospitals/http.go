package hospitals

import (
	"net/http"
	"peduli-covid/businesses/hospitals"
	controller "peduli-covid/controllers"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type HospitalController struct {
	hospitalUsecase hospitals.Usecase
}

func NewHospitalController(uc hospitals.Usecase) *HospitalController {
	return &HospitalController{
		hospitalUsecase: uc,
	}
}

func (ctrl *HospitalController) StoreFromAPI(c echo.Context) error {
	ctx := c.Request().Context()

	err := ctrl.hospitalUsecase.StoreFromAPI(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *HospitalController) FindByCityID(c echo.Context) error {
	ctx := c.Request().Context()

	cityID, _ := strconv.Atoi(c.QueryParam("city_id"))

	resp, err := ctrl.hospitalUsecase.FindByCityID(ctx, cityID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
