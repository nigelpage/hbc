package pennant

import (
	"net/http"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHeaderMenus() *[]common.HeaderMenu {
	hdrMenus := []common.HeaderMenu{
		*common.NewHeaderMenu("/pennant/:type", "pennant", http.MethodGet, handlers.CompetitionHandler),
		*common.NewHeaderMenu("/pennant/lock", "", http.MethodGet, handlers.LockAuthenticationHandler),
		*common.NewHeaderMenu("/pennant/authenticate", "", http.MethodPost, handlers.AuthenticationHandler),
	}
	return &hdrMenus
}