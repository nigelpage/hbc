package common

import (
	"strings"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

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
// N.B. If the Text field is empty the handler will be registered but not displayed in the menu
type HeaderMenu struct {
	Url string
	Text string
	Method string // Use values from net/http - http.MethodGet, http.MethodPost, etc...
	Handler echo.HandlerFunc
	SubMenus []HeaderMenu
}

var HeaderMenusAndSubMenus []HeaderMenu

func NewHeaderMenu(url string, text string, method string, handler echo.HandlerFunc) *HeaderMenu {
	if text != "" {
		text = strings.ToLower(text)
	}

	hdrMenu := HeaderMenu {
		Url: url,
		Text: text,
		Method: method,
		Handler: handler,
		SubMenus: []HeaderMenu{},
	}
	return &hdrMenu
}

func Pluralise(val int) string {
	if val == 1 {
		return ""
	}
	return "s"
}