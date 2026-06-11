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

func NewSeeder(db *sqlx.DB) *AppSeeder {
	return &AppSeeder{
		db:  db,
		log: logrus.New(),
	}
}

func Execute(db *sqlx.DB, log *logrus.Logger, seed string, total int) {
	seeder := NewSeeder(db)
	seeder.log = log
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

	defer func() {
		if err != nil {
			err = tx.Rollback()
			s.log.WithError(err).Error("failed to rollback transaction")
			return
		} else {
			err = tx.Commit()
			if err != nil {
				s.log.WithError(err).Error("failed to commit transaction")
			}
		}
	}()

	_, err = tx.Exec(`DELETE FROM examples`)
	if err != nil {
		s.log.WithError(err).Error("failed to delete examples")
		return
	}

	s.log.Info("examples table cleared successfully")
}
