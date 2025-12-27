package internal

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{
	urlPattern string
	function echo.HandlerFunc
}