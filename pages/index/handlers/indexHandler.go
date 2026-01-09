package handlers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/pages/index/templates"
)

func createMainPageFromTemplate(tickerMessages []string) templ.Component {
	return templates.BaseLayout(tickerMessages)
}

func IndexHandler(ctx echo.Context) error {
	tickerMessages := []string{
		"** PENNANT competition resumes on Saturday, so please make sure you attend training and/or skills & drills.",
		"** UPDATE this week's members draw and raffle has been postponed due to forecast heat.",
		"** INFO registration for 1 bowl singles competition closes at 5pm Friday.",
	}

	return templateRenderer(ctx, http.StatusOK, createMainPageFromTemplate(tickerMessages))
}