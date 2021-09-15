package reservations

import (
	"net/http"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/reservations"
	controller "peduli-covid/controllers"
	"peduli-covid/controllers/reservations/request"
	"strconv"

	echo "github.com/labstack/echo/v4"
)

type ReservationController struct {
	reservationUsecase reservations.Usecase
}

func NewReservationController(uc reservations.Usecase) *ReservationController {
	return &ReservationController{
		reservationUsecase: uc,
	}
}

func (ctrl *ReservationController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Reservations{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	req.UserID = middleware.GetUser(c).ID

	resp, err := ctrl.reservationUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *ReservationController) FindByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.reservationUsecase.FindByUserID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *ReservationController) FindByAdminID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.reservationUsecase.FindByAdminID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *ReservationController) UpdateStatus(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	req := request.Reservations{ID: id}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.reservationUsecase.UpdateStatusDone(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Data Updated")
}
