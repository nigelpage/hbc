package handlers

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/index/templates"
)

func createMainPageFromTemplate(tickerItems []*common.TickerItem) templ.Component {
	return templates.BaseLayout(tickerItems)
}

func IndexHandler(ctx echo.Context) error {
	tickerItems := []*common.TickerItem{
		common.NewTickerItem(
			time.Now(),
			time.Now().Add(time.Duration(7*24)*time.Hour),
			"pennant",
			"competition resumes on Saturday, so please make sure you attend training and/or skills & drills.",
		),
		common.NewTickerItem(
			time.Now(),
			time.Now().Add(time.Duration(7*24)*time.Hour),
			"update",
			"this week's members draw and raffle has been postponed due to forecast heat.",
		),
		common.NewTickerItem(
			time.Now(),
			time.Now().Add(time.Duration(7*24)*time.Hour),
			"info",
			"registration for 1 bowl singles competition closes at 5pm Friday.",
		),
	}

	return templateRenderer(ctx, http.StatusOK, createMainPageFromTemplate(tickerItems))
}