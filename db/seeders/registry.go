package seeders

import (
	"context"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type AppSeeder struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewSeeder(db *sqlx.DB, log *slog.Logger) *AppSeeder {
	return &AppSeeder{
		db:  db,
		log: log,
	}
}

func Execute(db *sqlx.DB, log *slog.Logger, seed string, total int) {
	seeder := NewSeeder(db, log)
	seeder.run(seed, total)
}

func (s *AppSeeder) run(seed string, total int) {
	switch seed {
	case "example":
		s.exampleSeed(total)
	case "role":
		s.roleSeed()
	case "admin":
		s.adminSeed()
	case "clear-example":
		s.clearExampleSeed()
	default:
		s.log.Warn("no seed to run", "seed", seed)
	}
}

func (s *AppSeeder) clearExampleSeed() {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		s.log.Error("failed to start transaction", "error", err)
		return
	}

	committed := false
	defer func() {
		if !committed {
			if rbErr := tx.Rollback(); rbErr != nil {
				s.log.Error("failed to rollback transaction", "error", rbErr)
			}
		}
	}()

	_, err = tx.Exec(`DELETE FROM examples`)
	if err != nil {
		s.log.Error("failed to delete examples", "error", err)
		return
	}

	if err = tx.Commit(); err != nil {
		s.log.Error("failed to commit transaction", "error", err)
		return
	}
	committed = true

	s.log.Info("examples table cleared successfully")
}
