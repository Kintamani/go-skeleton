package seeders

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/sirupsen/logrus"
)

func (s *AppSeeder) exampleSeed(total int) {
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
		logrus.Error("Error creating product categories")
	}

	logrus.Info("product_categories table seeded successfully")
}
