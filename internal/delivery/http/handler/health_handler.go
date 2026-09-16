package handler

import (
	"net/http"

	"github.com/Kintamani/go-skeleton/internal/delivery/http/response"
	"github.com/Kintamani/go-skeleton/internal/usecase"
	"github.com/labstack/echo/v4"
)

type HealthHandler struct {
	useCase *usecase.HealthUseCase
}

func NewHealthHandler(useCase *usecase.HealthUseCase) *HealthHandler {
	return &HealthHandler{useCase: useCase}
}

func (h *HealthHandler) Ping(c echo.Context) error {
	data := h.useCase.Ping(c.Request().Context())
	return response.JSON(c, http.StatusOK, "application is reachable", data)
}
