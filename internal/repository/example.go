package repository

import (
	"context"

	"github.com/Kintamani/go-skeleton/internal/entity"
)

type ExampleRepository interface {
	FindAll(ctx context.Context) ([]entity.Example, error)
}
