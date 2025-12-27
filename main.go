package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	i "github.com/nigelpage/hbc/internal"
	p "github.com/nigelpage/hbc/pages/pennant"
)

func registerHandlers(hdlrs []i.Handler, app *echo.Echo) error {
	/* Register handlers */
	for _, h := range hdlrs {
		switch h.Verb {
		case "GET":
			app.GET(h.UrlPattern, h.Function)
		case "POST":
			app.POST(h.UrlPattern, h.Function)
		case "PUT":
			app.PUT(h.UrlPattern, h.Function)
		case "DELETE":
			app.DELETE(h.UrlPattern, h.Function)
		case "PATCH":
			app.PATCH(h.UrlPattern, h.Function)
		case "HEAD":
			app.HEAD(h.UrlPattern, h.Function)
		case "OPTIONS":
			app.OPTIONS(h.UrlPattern, h.Function)
		// Invalid HTTP verb
		default:
			return fmt.Errorf("Invalid HTTP verb specified - %s - for url pattern - %s", h.Verb, h.UrlPattern)
		}
	}

	return nil
}

func main() {
	app := echo.New()
	/* Setup a handler for static files (e.g. CSS, JS etc...) */
	app.Static("/static", "/assets")
	
	/* Register HTTP handlers */
	
	/* ...for pennant page */

	err := registerHandlers(p.GetHandlers(), app)
	if err != nil {
		app.Logger.Fatal(err)	
	}

	/* Start HTTP server */
	app.Logger.Fatal(app.Start(":4000"))
}