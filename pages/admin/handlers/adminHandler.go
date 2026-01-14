package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/admin/templates"
)

func createAdminPageFromTemplate() templ.Component {
	return templates.PageLayout()
}

/* Main Admin page handler */
func AdminHandler(ctx echo.Context) error {
	return common.TemplateRenderer(ctx, http.StatusOK, createAdminPageFromTemplate())
}