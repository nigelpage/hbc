package common

import (
	"github.com/labstack/echo/v4"
)

// Hooks for security decisions
func IsAdminUser(ctx echo.Context) bool {
	return true;
}

func IsBowlsAdminUser(ctx echo.Context) bool {
	return true;
}