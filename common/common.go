package common

import (
	"fmt"
	"strings"

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

/* Substitute values into push URL if it contains parameters */
func SubstitutePushUrlParams(pushUrl string, substituteValues map[string]string) string {
	if !strings.Contains(pushUrl, ":") {
		return pushUrl
	}

var substitutedUrl string = ""

params := strings.Split(pushUrl, "/")
for _, param := range params[1:] {
		if len(param) > 0 && param[0] == ':' {
			substitutedUrl = fmt.Sprintf("%s/%s", substitutedUrl, substituteValues[param[1:]])
		} else {
			substitutedUrl = fmt.Sprintf("%s/%s", substitutedUrl, param)
		}
	}
	return substitutedUrl
}