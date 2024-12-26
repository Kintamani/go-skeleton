package bootstrap

import (
	"skeleton-services/api/routes"

	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type Bootstrap struct {
	DB        *sqlx.DB
	Framework *echo.Group
	Log       *logrus.Logger
	Validate  *validator.Validate
}

func App(b *Bootstrap) {
	routes := routes.Config{
		App: b.Framework,
	}
	routes.Api()
}
