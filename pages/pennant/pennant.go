package pennant

import (
	i "github.com/nigelpage/hbc/internal"
	h "github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHandlers() []i.Handler {
	return []i.Handler{
		{UrlPattern: "/pennant", 				Verb: "GET", 	Function: h.WeekendCompetitionHandler},
		{UrlPattern: "/pennant/:competition", 	Verb: "GET", 	Function: h.CompetitionHandler},
		{UrlPattern: "/pennant/authenticate", 	Verb: "POST", 	Function: h.AuthenticationHandler},
		{UrlPattern: "/pennant/lock", 			Verb: "GET", 	Function: h.LockAuthenticationHandler},
	}
}