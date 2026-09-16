package seeders

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AppSeeder struct {
	db  *sqlx.DB
	log *logrus.Logger
}

func NewSeeder(db *sqlx.DB, log *logrus.Logger) *AppSeeder {
	return &AppSeeder{
		db:  db,
		log: log,
	}
}

func Execute(db *sqlx.DB, log *logrus.Logger, seed string, total int) {
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
		s.log.WithField("seed", seed).Warn("no seed to run")
	}
}

func (s *AppSeeder) clearExampleSeed() {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		s.log.WithError(err).Error("failed to start transaction")
		return
	}

	committed := false
	defer func() {
		if !committed {
			if rbErr := tx.Rollback(); rbErr != nil {
				s.log.WithError(rbErr).Error("failed to rollback transaction")
			}
		}
	}()

	_, err = tx.Exec(`DELETE FROM examples`)
	if err != nil {
		s.log.WithError(err).Error("failed to delete examples")
		return
	}

	if err = tx.Commit(); err != nil {
		s.log.WithError(err).Error("failed to commit transaction")
		return
	}
	committed = true

	s.log.Info("examples table cleared successfully")
}
