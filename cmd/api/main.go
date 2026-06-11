package main

import (
	"github.com/Kintamani/go-skeleton/internal/app"
	"github.com/Kintamani/go-skeleton/internal/config"
)

func init() {
	config.ConfigEnv(
		config.WithPath("./"),
		config.WithFilename(".env"),
	).Initialize()
}

func main() {
	app.RunHTTP()
}
