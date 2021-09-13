package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
	Messages []string    `json:"messages"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Message = "Something not right"
	response.Messages = []string{err.Error()}

	return c.JSON(status, response)
}
