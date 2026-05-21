package database

import (
	"fmt"
	"skeleton-services/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func New(log *logrus.Logger) *sqlx.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s sslmode=%s TimeZone=Asia/Jakarta",
		config.ENV.DB.Username,
		config.ENV.DB.Password,
		config.ENV.DB.Host,
		config.ENV.DB.Port,
		config.ENV.DB.Database,
		config.ENV.DB.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.WithError(err).Fatal("failed to connect database")
	}

	db.SetMaxOpenConns(config.ENV.DB.MaxOpenCons)
	db.SetMaxIdleConns(config.ENV.DB.MaxIdleCons)
	db.SetConnMaxLifetime(time.Duration(config.ENV.DB.ConnMaxLifetime) * time.Second)

	if err := db.Ping(); err != nil {
		log.WithError(err).Fatal("failed to ping database")
	}

	return db
}
