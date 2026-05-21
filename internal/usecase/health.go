package usecase

import (
	"context"
	"skeleton-services/internal/model"
)

type HealthUseCase interface {
	Ping(ctx context.Context) model.HealthCheckResponse
}

type healthUseCase struct{}

func NewHealthUseCase() HealthUseCase {
	return &healthUseCase{}
}

func (u *healthUseCase) Ping(_ context.Context) model.HealthCheckResponse {
	return model.HealthCheckResponse{
		Message: "pong",
	}
}
