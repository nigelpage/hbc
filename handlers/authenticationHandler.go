package handlers

import (
	"net/http"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func AuthenticationHandler(ctx echo.Context) error {
	return templateRenderer(ctx, http.StatusOK, createMainPageFromTemplate("Weekend"))
}