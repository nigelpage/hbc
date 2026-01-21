package pennant

import (
	"net/http"

	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/pennant/:competition", "pennant", http.MethodGet, handlers.PennantHandler),
		*header.NewHeaderMenu("/pennant/lock", "", http.MethodGet, handlers.LockAuthenticationHandler),
		*header.NewHeaderMenu("/pennant/authenticate", "", http.MethodPost, handlers.AuthenticationHandler),
	}
	return &hdrMenus
}