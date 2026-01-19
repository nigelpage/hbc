package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	ct "github.com/nigelpage/hbc/common/templates"
	"github.com/nigelpage/hbc/pages/admin/templates"
)

/* Main Admin page handler */
func AdminHandler(ctx echo.Context) error {
	return common.TemplateRenderer(ctx, http.StatusOK, ct.CreatePageFromTemplate(ctx, templates.AdminLayout()))
}