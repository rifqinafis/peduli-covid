package bedtypes

import (
	"net/http"
	"peduli-covid/businesses/bedtypes"
	controller "peduli-covid/controllers"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type BedtypeController struct {
	bedtypeUsecase bedtypes.Usecase
}

func NewBedtypeController(uc bedtypes.Usecase) *BedtypeController {
	return &BedtypeController{
		bedtypeUsecase: uc,
	}
}

func (ctrl *BedtypeController) StoreFromAPI(c echo.Context) error {
	ctx := c.Request().Context()

	err := ctrl.bedtypeUsecase.StoreFromAPI(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *BedtypeController) FindByHospitalID(c echo.Context) error {
	ctx := c.Request().Context()

	hospitalID, _ := strconv.Atoi(c.QueryParam("hospital_id"))

	resp, err := ctrl.bedtypeUsecase.FindByHospitalID(ctx, hospitalID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
