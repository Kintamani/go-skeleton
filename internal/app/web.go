package app

import (
	"github.com/Kintamani/go-skeleton/internal/config"
	"github.com/Kintamani/go-skeleton/internal/delivery/http/handler"
	"github.com/Kintamani/go-skeleton/internal/delivery/http/route"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/database"
	httpserver "github.com/Kintamani/go-skeleton/internal/infrastructure/http"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/logger"
	"github.com/Kintamani/go-skeleton/internal/repository"
	"github.com/Kintamani/go-skeleton/internal/usecase"
)

func RunHTTP() {
	log := logger.New(config.ENV.App.Environment)
	app := httpserver.New()
	db := database.New(log)

	healthUseCase := usecase.NewHealthUseCase()
	exampleRepository := repository.NewExampleRepository(db)
	exampleUseCase := usecase.NewExampleUseCase(exampleRepository)

	routeConfig := route.Config{
		App:            app.Group("/api"),
		HealthHandler:  handler.NewHealthHandler(healthUseCase),
		ExampleHandler: handler.NewExampleHandler(exampleUseCase),
	}
	routeConfig.Setup()

	if err := app.Start(":" + config.ENV.App.Port); err != nil {
		log.WithError(err).Fatal("failed to start http server")
	}
}
