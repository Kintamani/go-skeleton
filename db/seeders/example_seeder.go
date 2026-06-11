package seeders

import (
	"github.com/brianvoe/gofakeit/v7"
)

func (s *AppSeeder) exampleSeed(total int) {
	if total <= 0 {
		total = 10
	}

	gofakeit.Seed(0)

	var (
		args  = make([]map[string]any, 0)
		query = "INSERT INTO examples (name) VALUES (:name)"
	)

	for i := 0; i < total; i++ {
		var (
			name = gofakeit.ProductCategory()
			arg  = make(map[string]any)
		)

		arg["name"] = name
		args = append(args, arg)
	}

	_, err := s.db.NamedExec(query, args)
	if err != nil {
		s.log.WithError(err).Error("failed to seed examples")
		return
	}

	s.log.WithField("total", len(args)).Info("examples table seeded successfully")
}
