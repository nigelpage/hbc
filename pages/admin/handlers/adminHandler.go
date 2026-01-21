package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	ht "github.com/nigelpage/hbc/pages/header/templates"
	"github.com/nigelpage/hbc/pages/admin/templates"
)

/* Main Admin page handler */
func AdminHandler(ctx echo.Context) error {
	if !common.IsAdminUser(ctx) {
		return ctx.NoContent(http.StatusForbidden)
		//return ctx.Redirect(http.StatusSeeOther, "/login")
	}
	return common.TemplateRenderer(ctx, http.StatusOK, ht.CreatePageFromTemplate(ctx, templates.AdminLayout()))
}