package internal

import (
	"github.com/labstack/echo/v4"
)

type Handler struct{
	UrlPattern string
	Verb       string
	Function echo.HandlerFunc
}