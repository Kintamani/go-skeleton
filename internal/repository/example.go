package repository

import (
	"context"
	"skeleton-services/internal/entity"
)

type ExampleRepository interface {
	FindAll(ctx context.Context) ([]entity.Example, error)
}
