package database

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/Kintamani/go-skeleton/internal/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(log *slog.Logger) *sqlx.DB {
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
		log.Error("failed to connect database", "error", err)
		os.Exit(1)
	}

	db.SetMaxOpenConns(config.ENV.DB.MaxOpenCons)
	db.SetMaxIdleConns(config.ENV.DB.MaxIdleCons)
	if config.ENV.DB.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(config.ENV.DB.ConnMaxLifetime) * time.Second)
	}

	return db
}
