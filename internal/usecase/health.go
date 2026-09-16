package usecase

import (
	"context"

	"github.com/Kintamani/go-skeleton/internal/model"
)

type HealthUseCase struct{}

func NewHealthUseCase() *HealthUseCase {
	return &HealthUseCase{}
}

func (u *HealthUseCase) Ping(_ context.Context) model.HealthCheckResponse {
	return model.HealthCheckResponse{
		Message: "pong",
	}
}
