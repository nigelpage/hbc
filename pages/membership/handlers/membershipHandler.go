package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	ht "github.com/nigelpage/hbc/pages/header/templates"
	"github.com/nigelpage/hbc/pages/membership/templates"
)

/* Main Membership page handler */
func MembershipHandler(ctx echo.Context) error {
	return common.TemplateRenderer(ctx, http.StatusOK, ht.CreatePageFromTemplate(ctx, templates.MembershipLayout()))
}