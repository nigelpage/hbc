package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"bufio"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/common/store/excel"
	"github.com/nigelpage/hbc/pages/index"
	"github.com/nigelpage/hbc/pages/pennant"
)

func registerHandlers(hdlrs []*common.Handler, app *echo.Echo) error {
	/* Register handlers */
	for _, h := range hdlrs {
		switch h.GetVerb() {
		case "GET":
			app.GET(h.GetUrlPattern(), h.GetFunction())
		case "POST":
			app.POST(h.GetUrlPattern(), h.GetFunction())
		case "PUT":
			app.PUT(h.GetUrlPattern(), h.GetFunction())
		case "DELETE":
			app.DELETE(h.GetUrlPattern(), h.GetFunction())
		case "PATCH":
			app.PATCH(h.GetUrlPattern(), h.GetFunction())
		case "HEAD":
			app.HEAD(h.GetUrlPattern(), h.GetFunction())
		case "OPTIONS":
			app.OPTIONS(h.GetUrlPattern(), h.GetFunction())
		// Invalid HTTP verb
		default:
			return fmt.Errorf("Invalid HTTP verb specified - %s - for url pattern - %s", h.GetVerb(), h.GetUrlPattern())
		}
	}

	return nil
}

func LoadEnvironmentVariables(filename string) error {
	// Setup environment variables
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("Failed to open environment variables file %w", err)
	}
	defer file.Close()

    // Create a new scanner to read the file line by line
    scanner := bufio.NewScanner(file)

    // Loop through the file and read each line
   for scanner.Scan() {
        line := scanner.Text() // Get the line as a string
		if strings.HasPrefix(line, "export ") {
			eVar := strings.Split(strings.TrimPrefix(line, "export "), "=")
			if len(eVar) != 2 {
				return fmt.Errorf("Invalid export value in environment variables file")
			}
			fmt.Printf("%s=%s\n", eVar[0], eVar[1])
			os.Setenv(eVar[0], eVar[1])
		}
    }

    // Check for errors during the scan
    if err := scanner.Err(); err != nil {
		return fmt.Errorf("Error reading environment variables file %w", err)
    }

	return nil
}

func pluraliseIfNotOne(val int) string {
	if val == 1 {
		return ""
	}
	return "s"
}

func main() {
	// Initialise database connection

	// Load environment variables from script file
	err := LoadEnvironmentVariables("./common/store/db/sqlc_cloud.sh")
	if err != nil {
		panic(err)
	}

	// Empty connection string forces use of environment variables
	ctx := context.Background();

	dbPool, err := pgxpool.New(ctx, "")
	if err != nil {
		panic("Unable to create connection pool")
	}

	defer dbPool.Close()

	// Migrate from Json to database
	// err = migrateFromJsonToDB(app.Pool, app.Queries)

	results, err := excel.UploadMembers(ctx, dbPool, "./common/store/excel/Members Draw list 02.01.2026.xlsx")
	fmt.Printf("%d member%s added, %d member%s updated\n", results.Added, pluraliseIfNotOne(results.Added),
														   results.Updated, pluraliseIfNotOne(results.Updated)) // ** Temporary

	// Initialise app
	app := common.NewApp(echo.New())
	app.Echo.Pre(middleware.RemoveTrailingSlash())
	
	// Setup a handler for static files (e.g. CSS, JS etc...)
	app.Echo.Static("/static", "pages")
	
	// Register HTTP handlers
	// ...for index page
	err = registerHandlers(index.GetHandlers(), app.Echo)
	if err != nil {
		app.Echo.Logger.Fatal(err)	
	}
	
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