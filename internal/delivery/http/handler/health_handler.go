package handler

import (
	"net/http"
	"skeleton-services/internal/delivery/http/response"
	"skeleton-services/internal/usecase"

	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	useCase usecase.HealthUseCase
}

func NewHealthHandler(useCase usecase.HealthUseCase) *HealthHandler {
	return &HealthHandler{useCase: useCase}
}

func (h *HealthHandler) Ping(c echo.Context) error {
	data := h.useCase.Ping(c.Request().Context())
	return response.JSON(c, http.StatusOK, "application is reachable", data)
}
