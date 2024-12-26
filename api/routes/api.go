package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Config struct {
	App *echo.Group
}

func (c *Config) Api() {
	v1 := c.App.Group("/v1")
	v1.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
}
