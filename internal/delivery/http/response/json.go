package response

import (
	"errors"
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
	if c.Response().Committed {
		return
	}

	code := http.StatusInternalServerError
	message := "Internal Server Error"

	var httpError *echo.HTTPError
	if errors.As(err, &httpError) {
		code = httpError.Code
		message = fmt.Sprint(httpError.Message)
	} else if err != nil {
		message = err.Error()
	}

	c.Logger().Error(err)
	_ = c.JSON(code, Envelope{
		Success: false,
		Status:  code,
		Message: message,
	})
}
