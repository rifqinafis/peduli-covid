package rsbedcovid

import (
	"net/http"
	"peduli-covid/businesses/rsbedcovid"
	controller "peduli-covid/controllers"

	"github.com/labstack/echo/v4"
)

type RSBedCovidController struct {
	rsbedcovid rsbedcovid.Usecase
}

func NewRSBedCovidController(uc rsbedcovid.Usecase) *RSBedCovidController {
	return &RSBedCovidController{
		rsbedcovid: uc,
	}
}

func (ctrl *RSBedCovidController) GetProvince(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.rsbedcovid.GetProvince(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *RSBedCovidController) GetCity(c echo.Context) error {
	ctx := c.Request().Context()

	provinceID := c.QueryParam("province_id")

	resp, err := ctrl.rsbedcovid.GetCity(ctx, provinceID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *RSBedCovidController) GetHospital(c echo.Context) error {
	ctx := c.Request().Context()

	provinceID := c.QueryParam("province_id")
	cityID := c.QueryParam("city_id")
	types := c.QueryParam("types")

	resp, err := ctrl.rsbedcovid.GetHospital(ctx, provinceID, cityID, types)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *RSBedCovidController) GetBedDetail(c echo.Context) error {
	ctx := c.Request().Context()

	hospitalID := c.QueryParam("hospital_id")
	types := c.QueryParam("types")

	resp, err := ctrl.rsbedcovid.GetBedDetail(ctx, hospitalID, types)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *RSBedCovidController) GetHospitalLocation(c echo.Context) error {
	ctx := c.Request().Context()

	hospitalID := c.QueryParam("hospital_id")

	resp, err := ctrl.rsbedcovid.GetHospitalLocation(ctx, hospitalID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
