package payments

import (
	"net/http"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/payments"
	controller "peduli-covid/controllers"
	"peduli-covid/controllers/payments/request"

	echo "github.com/labstack/echo/v4"
)

type PaymentController struct {
	paymentUsecase payments.Usecase
}

func NewPaymentController(uc payments.Usecase) *PaymentController {
	return &PaymentController{
		paymentUsecase: uc,
	}
}

func (ctrl *PaymentController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Payments{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	userID := middleware.GetUser(c).ID

	err := ctrl.paymentUsecase.Store(ctx, userID, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *PaymentController) FindByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.paymentUsecase.FindByUserID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
