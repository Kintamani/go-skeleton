package main

import (
	"skeleton-services/adapters"
	"skeleton-services/bootstrap"
	"skeleton-services/config"
)

func init() {
	config.ConfigEnv(
		config.WithPath("./"),
		config.WithFilename(".env"),
	).Initialize()
}

func main() {
	// cmd.Execute()

	log := adapters.NewLogger(config.ENV.App.Stage)
	validate := adapters.NewValidator()
	app := adapters.NewEcho(log)
	db := adapters.NewDatabase(log)
	group := app.Group("/api")

	bootstrap.App(&bootstrap.Bootstrap{
		DB:        db,
		Framework: group,
		Log:       log,
		Validate:  validate,
	})

	port := config.ENV.App.Port
	app.Logger.Fatal(app.Start(":" + port))
}
