package app

import (
	"skeleton-services/config"
	"skeleton-services/internal/delivery/http/handler"
	"skeleton-services/internal/delivery/http/route"
	"skeleton-services/internal/infrastructure/database"
	httpserver "skeleton-services/internal/infrastructure/http"
	"skeleton-services/internal/infrastructure/logger"
	"skeleton-services/internal/infrastructure/persistence"
	"skeleton-services/internal/usecase"
)

func RunHTTP() {
	log := logger.New(config.ENV.App.Environment)
	app := httpserver.New()
	db := database.New(log)

	healthUseCase := usecase.NewHealthUseCase()
	exampleRepository := persistence.NewExampleRepository(db)
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
