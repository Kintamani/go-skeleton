package app

import (
	"os"

	"github.com/Kintamani/go-skeleton/internal/config"
	"github.com/Kintamani/go-skeleton/internal/delivery/http/handler"
	"github.com/Kintamani/go-skeleton/internal/delivery/http/response"
	"github.com/Kintamani/go-skeleton/internal/delivery/http/route"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/database"
	"github.com/Kintamani/go-skeleton/internal/infrastructure/logger"
	"github.com/Kintamani/go-skeleton/internal/repository"
	"github.com/Kintamani/go-skeleton/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RunHTTP() {
	log := logger.New(config.ENV.App.Environment)
	db := database.New(log)
	defer db.Close()

	app := echo.New()
	app.HideBanner = true
	app.HTTPErrorHandler = response.HTTPErrorHandler
	app.Use(middleware.Recover())

	exampleRepository := repository.NewExampleRepository(db)
	exampleUseCase := usecase.NewExampleUseCase(exampleRepository)

	routeConfig := route.Config{
		App:            app.Group("/api"),
		HealthHandler:  handler.NewHealthHandler(),
		ExampleHandler: handler.NewExampleHandler(exampleUseCase),
	}
	routeConfig.Setup()

	if err := app.Start(":" + config.ENV.App.Port); err != nil {
		log.Error("failed to start http server", "error", err)
		os.Exit(1)
	}
}
