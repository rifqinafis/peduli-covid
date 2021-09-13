package cities

import (
	"net/http"
	"peduli-covid/businesses/cities"
	controller "peduli-covid/controllers"
	"peduli-covid/controllers/cities/request"

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

func (ctrl *CityController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Cities{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.cityUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
