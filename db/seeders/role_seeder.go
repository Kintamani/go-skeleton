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
		s.log.WithError(err).Error("failed to seed roles")
		return
	}

	s.log.WithField("total", len(args)).Info("roles table seeded successfully")
}
