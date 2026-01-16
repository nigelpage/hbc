package handlers

import (
	"net/http"
	"time"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	ct "github.com/nigelpage/hbc/common/templates")

func createPageFromTemplate(tickerItems *[]common.TickerItem) templ.Component {
	return ct.HeaderLayout(tickerItems)
}

func IndexHandler(ctx echo.Context) error {
	tickerItems := []common.TickerItem {
		{
			StartAt: time.Now(),
			EndAt: time.Now().Add(time.Duration(7*24)*time.Hour),
			Category: "pennant",
			Message: "competition resumes on Saturday, so please make sure you attend training and/or skills & drills.",
		},
		{
			StartAt: time.Now(),
			EndAt: time.Now().Add(time.Duration(7*24)*time.Hour),
			Category: "update",
			Message: "this week's members draw and raffle has been postponed due to forecast heat.",
		},
		{
			StartAt: time.Now(),
			EndAt: time.Now().Add(time.Duration(7*24)*time.Hour),
			Category: "info",
			Message: "registration for 1 bowl singles competition closes at 5pm Friday.",
		},
	}

	return common.TemplateRenderer(ctx, http.StatusOK, createPageFromTemplate(&tickerItems))
}