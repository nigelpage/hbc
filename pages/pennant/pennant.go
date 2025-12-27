package pennant

import (
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/internal/handlers"
)

func GetHandlers() []handlers.Handler {
	return []handlers.Handler{
		{"/pennant/", 				handlers.WeekendCompetitionHandler},
		{"/pennant/:competition", 	handlers.CompetitionHandler},
		{"/pennant/authenticate", 	handlers.AuthenticationHandler},
		{"/pennant/lock", 			handlers.LockAuthenticationHandler},
	}
}