package adapters

import (
	"skeleton-services/api/responses"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func NewEcho(log *logrus.Logger) *echo.Echo {
	e := echo.New()
	// e.HideBanner = true
	e.HTTPErrorHandler = responses.HTTPErrorHandler
	// e.Logger.SetLevel(log.InfoLevel)
	return e
}
