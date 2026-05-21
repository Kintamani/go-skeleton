package response

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Envelope struct {
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(c echo.Context, status int, message string, data any) error {
	return c.JSON(status, Envelope{
		Success: true,
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func HTTPErrorHandler(err error, c echo.Context) {
	httpError, ok := err.(*echo.HTTPError)
	if !ok {
		httpError = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Error(httpError.Message)
	_ = c.JSON(httpError.Code, Envelope{
		Success: false,
		Status:  httpError.Code,
		Message: fmt.Sprint(httpError.Message),
	})
}
