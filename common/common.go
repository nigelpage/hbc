package common

import (
	"fmt"
	"strings"
	"time"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

var HeaderMenusAndHandlers = map[string]*[]HeaderMenuAndHandler{}

// Check to see if the base page is already loaded
func IsHeaderLoaded(ctx echo.Context) (bool, error) {
	name := "X-HBC-HeaderLoaded"
	isLoaded := ctx.Request().Header.Get(name)
	if isLoaded != "" {
		torf, err := strconv.ParseBool(isLoaded)
		if err != nil {			
			return false, fmt.Errorf("Invalid value for header '%s' - must be 'true' or 'false', not '%s'", name, isLoaded)
		}
		return torf, nil
	}
	return false, nil
}

/* Template renderer */
func TemplateRenderer(ctx echo.Context, statusCode int, cmp templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	
	if err := cmp.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}

// Ticker items

type TickerItem struct {
	StartAt		time.Time
	EndAt		time.Time
	Category	string
	Message		string
}

// Menus and handlers

// The header can have one or two levels of menu - e.g. /home or /bowls/pennant
// The Url field is relative to the location of the web site - e.g. https://www.heathmontbowlsclub.com.au/bowls/pennant
// The Text field is the text that will be displayed in the menu and it will be forced to be all lowercase
// Each page wishing to support one or more menu items must have a 'func registerMenuItems([]HeaderMenuItem hmi)'
// N.B. If the Text field is empty the handler will be registered but not displayed in the menu
type HeaderMenuAndHandler struct {
	Url string
	IsDropDown bool
	Text string
	Method string
	Handler echo.HandlerFunc
}

func ValidateHeaderMenusAndHandlers(hmahs []HeaderMenuAndHandler) error {
	var l1 string
	for _, hmah := range hmahs {
		if hmah.Url != "" && len(hmah.Url) > 1 {
			// Make sure the first character of the url is a '/'
			if hmah.Url[0] != '/' {
				hmah.Url = "/" + hmah.Url
			}
			parts := strings.Split(hmah.Url, "/")

			// Because the Url starts with a '/', the first part will always be empty
			if len(parts) < 2 || len(parts) > 3 {
				return fmt.Errorf("Invalid menu item url - %s - must be of the form /level1 or /level1/level2", hmah.Url)
			}

			if l1 != "" && l1 != parts[1] {
				return fmt.Errorf("Different level 1 items found - %s and %s - only one level 1 menu item per page is allowed", l1, parts[1])
			}

			l1 = parts[1]
		} else {
			if hmah.Url != "" {
				return fmt.Errorf("Cannot associate a URL with a drop down menu header - %s", hmah.Text)
			}
		}

		// Force the menu text to be all lowercase
		// Blank text means it will not appear in the menu but the handler is still registered
		if len(hmah.Text) > 0 {
			hmah.Text = strings.ToLower(hmah.Text)
		}

		if hmah.Method == "" || len(hmah.Method) == 0 {
			return fmt.Errorf("No HTTP method specified for menu item url - %s", hmah.Url)
		}
		hmah.Method = strings.ToUpper(hmah.Method)

		if hmah.Handler == nil {
			return fmt.Errorf("No handler function specified for menu item url - %s", hmah.Url)
		}
	}

	return nil
}


func Pluralise(val int) string {
	if val == 1 {
		return ""
	}
	return "s"
}