package usecase

import (
	"context"

	"github.com/Kintamani/go-skeleton/internal/model"
	"github.com/Kintamani/go-skeleton/internal/repository"
)

type ExampleUseCase struct {
	exampleRepository *repository.ExampleRepository
}

func NewExampleUseCase(exampleRepository *repository.ExampleRepository) *ExampleUseCase {
	return &ExampleUseCase{exampleRepository: exampleRepository}
}

func (u *ExampleUseCase) FindAll(ctx context.Context) ([]model.ExampleResponse, error) {
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
