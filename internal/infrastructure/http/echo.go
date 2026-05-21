package httpserver

import (
	"skeleton-services/internal/delivery/http/response"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	app := echo.New()
	app.HTTPErrorHandler = response.HTTPErrorHandler
	return app
}
