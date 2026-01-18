package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/index"
	"github.com/nigelpage/hbc/pages/pennant"
)

func registerHeaderMenus(category string, isSubMenu bool, hdrMenus *[]common.HeaderMenu, app *echo.Echo) (*[]common.HeaderMenu, error) {
	if (hdrMenus != nil) {
		for _, menu := range *hdrMenus {
			url := menu.Url
			if isSubMenu {
				url = fmt.Sprintf("/%s%s", category, menu.Url)
			}

			switch menu.Method {
				case http.MethodGet:
					app.GET(url, menu.Handler)
				case http.MethodPost:
					app.POST(url, menu.Handler)
				case http.MethodPut:
					app.PUT(url, menu.Handler)
				case http.MethodDelete:
					app.DELETE(url, menu.Handler)
				case http.MethodPatch:
					app.PATCH(url, menu.Handler)
				case http.MethodHead:
					app.HEAD(url, menu.Handler)
				case http.MethodOptions:
					app.OPTIONS(url, menu.Handler)
				// Invalid HTTP method
				default:
					return nil, fmt.Errorf("Invalid HTTP method specified - %s - for url pattern - %s", menu.Method, url)
				}
		}
	}
	return hdrMenus, nil
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

	// N.B. Just used whilst developing
	whileDev(ctx, dbPool)

	// Initialise app
	app := common.NewApp(echo.New())
	app.Echo.Pre(middleware.RemoveTrailingSlash())
	
	// Setup a handler for static files (e.g. CSS, JS etc...)
	app.Echo.Static("/static", "pages")
	
	// Register HTTP handlers and menus
	var cat string

	// ...for index page
	cat = "home"
	menus, err := registerHeaderMenus(cat, false, index.GetHeaderMenus(), app.Echo)
	if err != nil {
		app.Echo.Logger.Fatal(err)	
	}

	// Add to the collection of all menus and handlers
	common.HeaderMenusAndSubMenus = *menus
	
	// ...for bowls pages
	
	cat = "bowls"
	menus, err = registerHeaderMenus(cat, false, nil, nil)
	if err != nil {
		app.Echo.Logger.Fatal(err)	
	}
	menuContainer := *common.NewHeaderMenu("", cat, "", nil)
	
	subMenus, err := registerHeaderMenus(cat, true, pennant.GetHeaderMenus(), app.Echo)
	if err != nil {
		app.Echo.Logger.Fatal(err)	
	}
	menuContainer.SubMenus = append(menuContainer.SubMenus, *subMenus...)
	common.HeaderMenusAndSubMenus = append(common.HeaderMenusAndSubMenus, menuContainer)

	// Setup logging middleware
	app.Echo.Use(middleware.RequestLogger())

	// Start HTTP server
	app.Echo.Logger.Fatal(app.Echo.Start(":4000"))
}