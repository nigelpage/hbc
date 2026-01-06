package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	i "github.com/nigelpage/hbc/internal"
	"github.com/nigelpage/hbc/pages/pennant"
	"github.com/nigelpage/hbc/pages/index"
	dbstore "github.com/nigelpage/hbc/store/db"
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
	// Initialise database connection
	// Forces use of environment variables DBNAME, PGUSER and PGPASSWORD
	connString := ""

	pool, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	// Initialise app
	app := NewApp(echo.New(), pool, dbstore.New(pool))

	// Migrate from Json to database
	// err = migrateFromJsonToDB(app.Pool, app.Queries)
	// if err != nil {
	// 	app.Echo.Logger.Fatal(err)	
	// }

	app.Echo.Pre(middleware.RemoveTrailingSlash())
	
	// Setup a handler for static files (e.g. CSS, JS etc...)
	app.Echo.Static("/static", "pages")
	
	// Register HTTP handlers
	// ...for index page
	err = registerHandlers(index.GetHandlers(), app.Echo)
	
	// ...for pennant page

	err = registerHandlers(pennant.GetHandlers(), app.Echo)
	if err != nil {
		app.Echo.Logger.Fatal(err)	
	}

	// Setup logging middleware
	app.Echo.Use(middleware.RequestLogger())

	// Start HTTP server
	app.Echo.Logger.Fatal(app.Echo.Start(":4000"))
}