package notifications

import (
	"net/http"
	"peduli-covid/app/middleware"
	"peduli-covid/businesses/notifications"
	controller "peduli-covid/controllers"

	echo "github.com/labstack/echo/v4"
)

type NotificationController struct {
	notificationUsecase notifications.Usecase
}

func NewNotificationController(uc notifications.Usecase) *NotificationController {
	return &NotificationController{
		notificationUsecase: uc,
	}
}

func (ctrl *NotificationController) FindByUserID(c echo.Context) error {
	ctx := c.Request().Context()
	userID := middleware.GetUser(c).ID

	resp, err := ctrl.notificationUsecase.FindByUserID(ctx, userID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, resp)
}
