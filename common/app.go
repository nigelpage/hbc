package common

import (
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo	*echo.Echo
}

func NewApp(echo *echo.Echo) (*App) {
	// Initialise Echo web server
	app := &App {
		Echo:	echo,
	}

	return app
}