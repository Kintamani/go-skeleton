package seeders

func (s *AppSeeder) roleSeed() {
	query := `
		INSERT INTO roles (name)
		VALUES (:name)
		ON CONFLICT (name) DO NOTHING
	`

	args := []map[string]any{
		{"name": "admin"},
		{"name": "user"},
	}

	_, err := s.db.NamedExec(query, args)
	if err != nil {
		s.log.Error("failed to seed roles", "error", err)
		return
	}

	s.log.Info("roles table seeded successfully", "total", len(args))
}
