package pennant

import (
	i "github.com/nigelpage/hbc/internal"
	h "github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHandlers() []i.Handler {
	return []i.Handler{
		{UrlPattern: "/pennant/", 				Function: h.WeekendCompetitionHandler},
		{UrlPattern: "/pennant/:competition", 	Function: h.CompetitionHandler},
		{UrlPattern: "/pennant/authenticate", 	Function: h.AuthenticationHandler},
		{UrlPattern: "/pennant/lock", 			Function: h.LockAuthenticationHandler},
	}
}