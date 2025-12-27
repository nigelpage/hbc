package main

import (
	"github.com/labstack/echo/v4"
	// "github.com/nigelpage/hbc/internal"
	"github.com/nigelpage/hbc/pages/pennant"
)

func main() {
	app := echo.New()
	/* Setup a handler for static files (e.g. CSS, JS etc...) */
	app.Static("/static", "./assets")
	
	/* Register handlers */
	for _, h := range pennant.GetHandlers() {
		app.GET(h.Handler.urlPattern, h.function)
	}

	/* Start HTTP server */
	app.Logger.Fatal(app.Start(":4000"))
}