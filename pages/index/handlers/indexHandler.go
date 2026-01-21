package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	ht "github.com/nigelpage/hbc/pages/header/templates"
	"github.com/nigelpage/hbc/pages/index/templates"
)

func IndexHandler(ctx echo.Context) error {

	return common.TemplateRenderer(ctx, http.StatusOK, ht.CreatePageFromTemplate(ctx, templates.IndexLayout()))
}