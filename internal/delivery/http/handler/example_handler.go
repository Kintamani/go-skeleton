package handler

import (
	"net/http"

	"github.com/Kintamani/go-skeleton/internal/delivery/http/response"
	"github.com/Kintamani/go-skeleton/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ExampleHandler struct {
	useCase *usecase.ExampleUseCase
}

func NewExampleHandler(useCase *usecase.ExampleUseCase) *ExampleHandler {
	return &ExampleHandler{useCase: useCase}
}

func (h *ExampleHandler) FindAll(c echo.Context) error {
	data, err := h.useCase.FindAll(c.Request().Context())
	if err != nil {
		return err
	}

	return response.JSON(c, http.StatusOK, "examples fetched", data)
}
