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
		db: db,
	}
}

func Execute(db *sqlx.DB, table string, total int) {
	seed := NewSeeder(db)
	seed.run(table, total)
}

func (s *AppSeeder) run(table string, total int) {
	switch table {
	case "example":
		s.exampleSeed(total)
	case "delete-all":
		s.deleteAllSeed()
	default:
		s.log.Warn("No seed to run")
	}

	//	if table != "" {
	//		s.log.Info("Seed ran successfully")
	//		log.Info().Msg("Exiting ...")
	//		if err := adapter.Adapters.Unsync(); err != nil {
	//			s.log.Error("Error while closing database connection")
	//		}
	//		os.Exit(0)
	//	}
}

func (s *AppSeeder) deleteAllSeed() {
	tx, err := s.db.BeginTxx(context.Background(), nil)
	if err != nil {
		logrus.Error("Error starting transaction")
		return
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
			logrus.Error("Error rolling back transaction")
			return
		} else {
			err = tx.Commit()
			if err != nil {
				logrus.Error("Error committing transaction")
			}
		}
	}()

	_, err = tx.Exec(`DELETE FROM users`)
	if err != nil {
		logrus.Error("Error deleting users")
		return
	}

	logrus.Info("users table deleted successfully")

	_, err = tx.Exec(`DELETE FROM roles`)
	if err != nil {
		logrus.Error("Error deleting roles")
		return
	}
	logrus.Info("roles table deleted successfully")

	logrus.Info("=== All tables deleted successfully ===")
}
