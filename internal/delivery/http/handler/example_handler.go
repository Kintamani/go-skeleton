package handler

import (
	"net/http"
	"skeleton-services/internal/delivery/http/response"
	"skeleton-services/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	useCase usecase.ExampleUseCase
}

func NewExampleHandler(useCase usecase.ExampleUseCase) *ExampleHandler {
	return &ExampleHandler{useCase: useCase}
}

func (h *ExampleHandler) FindAll(c echo.Context) error {
	data, err := h.useCase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return response.JSON(c, http.StatusOK, "examples fetched", data)
}
