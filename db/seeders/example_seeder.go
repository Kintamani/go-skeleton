package seeders

import "fmt"

func (s *AppSeeder) exampleSeed(total int) {
	if total <= 0 {
		total = 10
	}

	args := make([]map[string]any, total)
	for i := 0; i < total; i++ {
		args[i] = map[string]any{
			"name": fmt.Sprintf("Example Item %d", i+1),
		}
	}

	_, err := s.db.NamedExec("INSERT INTO examples (name) VALUES (:name)", args)
	if err != nil {
		s.log.Error("failed to seed examples", "error", err)
		return
	}

	s.log.Info("examples table seeded successfully", "total", len(args))
}
