package main

import (
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/pennant/handlers"
)

type handler struct{
	urlPattern string
	function echo.HandlerFunc
}

var handlersList = []handler{
	{"/", handlers.WeekendCompetitionHandler},
	{"/:competition", handlers.CompetitionHandler},
	{"/authenticate", handlers.AuthenticationHandler},
	{"/lock", handlers.LockAuthenticationHandler},
}

func main() {
	app := echo.New()
	/* Setup a handler for static files (e.g. CSS, JS etc...) */
	app.Static("/static", "./assets")
	
	/* Register handlers */
	for _, h := range handlersList {
		app.GET(h.urlPattern, h.function)
	}

	/* Start HTTP server */
	app.Logger.Fatal(app.Start(":4000"))
}