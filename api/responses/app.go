package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Default struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}

type Response struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func HTTPErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if !ok {
		he = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.Logger().Error(he.Message)
	c.JSONPretty(he.Code, Default{
		Success: false,
		Status:  he.Code,
		Message: he.Message,
	}, " ")
}

func BasesResponses(success bool, status int, messages, data interface{}) *Response {
	r := Response{
		Success: success,
		Status:  status,
		Message: messages,
		Data:    data,
	}
	return &r
}
