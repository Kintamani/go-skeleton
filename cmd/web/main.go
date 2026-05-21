package main

import (
	"skeleton-services/config"
	"skeleton-services/internal/app"
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
