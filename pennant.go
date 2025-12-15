package main

import (
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/pennant/handlers"
)

func main() {
	app := echo.New()
	/* Setup a handler for static files (e.g. CSS, JS etc...) */
	app.Static("/static", "./assets")
	
	/* Setup main handler */
	app.GET("/", handlers.WeekendCompetitionHandler)
	app.GET("/:competition", handlers.CompetitionHandler)
	app.POST("/authenticate", handlers.AuthenticationHandler)
	app.GET("/lock", handlers.LockAuthenticationHandler)

	/* Start HTTP server */
	app.Logger.Fatal(app.Start(":4000"))
}