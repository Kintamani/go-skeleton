package usecase

import (
	"context"
	"skeleton-services/internal/model"
	"skeleton-services/internal/repository"
)

type ExampleUseCase interface {
	FindAll(ctx context.Context) ([]model.ExampleResponse, error)
}

type exampleUseCase struct {
	exampleRepository repository.ExampleRepository
}

func NewExampleUseCase(exampleRepository repository.ExampleRepository) ExampleUseCase {
	return &exampleUseCase{exampleRepository: exampleRepository}
}

func (u *exampleUseCase) FindAll(ctx context.Context) ([]model.ExampleResponse, error) {
	examples, err := u.exampleRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]model.ExampleResponse, 0, len(examples))
	for _, example := range examples {
		responses = append(responses, model.ExampleResponse{
			ID:        example.ID,
			Name:      example.Name,
			CreatedAt: example.CreatedAt,
		})
	}

	return responses, nil
}
