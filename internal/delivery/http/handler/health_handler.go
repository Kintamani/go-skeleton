package handler

import (
	"net/http"

	"github.com/Kintamani/go-skeleton/internal/delivery/http/response"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Ping(c echo.Context) error {
	return response.JSON(c, http.StatusOK, "application is reachable", map[string]string{
		"message": "pong",
	})
}
