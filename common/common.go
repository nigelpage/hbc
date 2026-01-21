package common

import (
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

func Pluralise(val int) string {
	if val == 1 {
		return ""
	}
	return "s"
}