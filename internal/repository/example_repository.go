package repository

import (
	"context"

	"github.com/Kintamani/go-skeleton/internal/entity"
	"github.com/jmoiron/sqlx"
)

type exampleRepository struct {
	db *sqlx.DB
}

func NewExampleRepository(db *sqlx.DB) ExampleRepository {
	return &exampleRepository{db: db}
}

func (r *exampleRepository) FindAll(ctx context.Context) ([]entity.Example, error) {
	examples := make([]entity.Example, 0)
	err := r.db.SelectContext(ctx, &examples, "SELECT id, name, created_at FROM examples ORDER BY id ASC")
	if err != nil {
		return nil, err
	}

	return examples, nil
}
