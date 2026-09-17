package seeders

import "golang.org/x/crypto/bcrypt"

func (s *AppSeeder) adminSeed() {
	s.roleSeed()

	var roleID int64
	err := s.db.Get(&roleID, `SELECT id FROM roles WHERE name = $1 LIMIT 1`, "admin")
	if err != nil {
		s.log.Error("failed to find admin role", "error", err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		s.log.Error("failed to hash admin password", "error", err)
		return
	}

	query := `
		INSERT INTO users (role_id, name, email, password)
		VALUES (:role_id, :name, :email, :password)
		ON CONFLICT (email) DO NOTHING
	`

	args := map[string]any{
		"role_id":  roleID,
		"name":     "Administrator",
		"email":    "admin@example.com",
		"password": string(hashedPassword),
	}

	_, err = s.db.NamedExec(query, args)
	if err != nil {
		s.log.Error("failed to seed admin user", "error", err)
		return
	}

	s.log.Info("admin user seeded successfully", "email", "admin@example.com")
}
