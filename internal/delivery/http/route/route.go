package route

import (
	"skeleton-services/internal/delivery/http/handler"

	"github.com/labstack/echo/v4"
)

type Config struct {
	App            *echo.Group
	HealthHandler  *handler.HealthHandler
	ExampleHandler *handler.ExampleHandler
}

func (c *Config) Setup() {
	v1 := c.App.Group("/v1")
	v1.GET("/ping", c.HealthHandler.Ping)
	v1.GET("/examples", c.ExampleHandler.FindAll)
}
