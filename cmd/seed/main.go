package main

import (
	"flag"

	"github.com/Kintamani/go-skeleton/db/seeders"
	"github.com/Kintamani/go-skeleton/internal/config"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/database"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/logger"
)

func init() {
	config.Load()
}

func main() {
	seed := flag.String("seed", "example", "seed target name")
	total := flag.Int("total", 10, "total records to generate")
	flag.Parse()

	log := logger.New(config.ENV.App.Environment)
	db := database.New(log)
	defer db.Close()

	seeders.Execute(db, log, *seed, *total)
}
