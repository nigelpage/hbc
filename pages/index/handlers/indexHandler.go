package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/pages/index/templates"
)

func createMainPageFromTemplate() templ.Component {
	return templates.BaseLayout()
}

func IndexHandler(ctx echo.Context) error {
	return templateRenderer(ctx, http.StatusOK, createMainPageFromTemplate())
}