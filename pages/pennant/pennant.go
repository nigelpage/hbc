package pennant

import (
	"net/http"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHeaderMenusAndHandlers() *[]common.HeaderMenuAndHandler {
	hmahs := []common.HeaderMenuAndHandler{
		{Url: "/pennant", Text: "pennant", Method: http.MethodGet, Handler: handlers.WeekendCompetitionHandler},
		{Url: "/pennant/authenticate", Text: "", Method: http.MethodPost, Handler: handlers.AuthenticationHandler},
		{Url: "/pennant/lock", Text: "", Method: http.MethodGet, Handler: handlers.LockAuthenticationHandler},
	}
	return &hmahs
}