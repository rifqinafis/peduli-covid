package cities

import (
	"net/http"
	"peduli-covid/businesses/cities"
	controller "peduli-covid/controllers"

	echo "github.com/labstack/echo/v4"
)

type CityController struct {
	cityUsecase cities.Usecase
}

func NewCityController(uc cities.Usecase) *CityController {
	return &CityController{
		cityUsecase: uc,
	}
}

func (ctrl *CityController) StoreFromAPI(c echo.Context) error {
	ctx := c.Request().Context()

	err := ctrl.cityUsecase.StoreFromAPI(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *CityController) FindByProvinceCode(c echo.Context) error {
	ctx := c.Request().Context()

	provinceCode := c.QueryParam("province_code")

	resp, err := ctrl.cityUsecase.FindByProvinceCode(ctx, provinceCode)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
