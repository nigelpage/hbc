package common

import (
	"time"
	"strings"

	"github.com/labstack/echo/v4"
)

// Page handlers

type Handler struct{
	urlPattern	string
	verb		string
	function	echo.HandlerFunc
}

func NewHandler(urlPattern string, verb string, function echo.HandlerFunc) *Handler {
	return &Handler{
		urlPattern: urlPattern,
		verb:       verb,
		function:   function,
	}
}

func (h *Handler) GetUrlPattern() string {
	return h.urlPattern
}

func (h *Handler) GetVerb() string {
	return h.verb
}

func (h *Handler) GetFunction() echo.HandlerFunc {
	return h.function
}

// End of page handlers
// Ticker items

type TickerItem struct {
	startAt		time.Time
	endAt		time.Time
	category	string
	message		string
}

func NewTickerItem(startAt time.Time, endAt time.Time, category string, message string) *TickerItem {
	return &TickerItem{
		startAt:	startAt,
		endAt:		endAt,
		category:	strings.ToUpper(category),
		message:	message,
	}
}

func (ti *TickerItem) GetStartAt() time.Time {
	return ti.startAt
}

func (ti *TickerItem) GetEndAt() time.Time {
	return ti.endAt
}

func (ti *TickerItem) GetCategory() string {
	return ti.category
}

func (ti *TickerItem) GetMessage() string {
	return ti.message
}