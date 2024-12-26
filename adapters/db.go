package adapters

import (
	"fmt"
	"skeleton-services/config"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// buat koneksi database
func NewDatabase(log *logrus.Logger) *sqlx.DB {
	username := config.ENV.DB.Username
	password := config.ENV.DB.Password
	host := config.ENV.DB.Host
	port := config.ENV.DB.Port
	database := config.ENV.DB.Database
	sslmode := config.ENV.DB.SSLMode

	maxpool := config.ENV.DB.MaxOpenCons
	maxidleconn := config.ENV.DB.MaxIdleCons
	maxconnlifetime := config.ENV.DB.ConnMaxLifetime

	dns := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=%s TimeZone=Asia/Jakarta", username, password, host, port, database, sslmode)
	db, err := sqlx.Connect("postgres", dns)
	if err != nil {
		log.Fatal("ERROR: $s", err)
	}

	db.SetMaxOpenConns(maxpool)
	db.SetMaxIdleConns(maxidleconn)
	db.SetConnMaxLifetime(time.Duration(maxconnlifetime))

	err = db.Ping()
	if err != nil {
		log.Fatal("ERROR: $s", err)
	}

	return db
}
