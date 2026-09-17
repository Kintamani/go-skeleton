package usecase

import (
	"context"

	"github.com/Kintamani/go-skeleton/internal/entity"
	"github.com/Kintamani/go-skeleton/internal/repository"
)

type ExampleUseCase struct {
	repo *repository.ExampleRepository
}

func NewExampleUseCase(repo *repository.ExampleRepository) *ExampleUseCase {
	return &ExampleUseCase{repo: repo}
}

func (u *ExampleUseCase) FindAll(ctx context.Context) ([]entity.Example, error) {
	return u.repo.FindAll(ctx)
}
