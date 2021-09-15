package invoices

import (
	"net/http"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/invoices"
	controller "peduli-covid/controllers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type InvoiceController struct {
	invoiceUsecase invoices.Usecase
}

func NewInvoiceController(uc invoices.Usecase) *InvoiceController {
	return &InvoiceController{
		invoiceUsecase: uc,
	}
}

func (ctrl *InvoiceController) FindByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.invoiceUsecase.FindByUserID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}

func (ctrl *InvoiceController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()
	id, _ := strconv.Atoi(c.Param("id"))

	resp, err := ctrl.invoiceUsecase.GetByID(ctx, id)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
