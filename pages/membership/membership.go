package membership

import (
	"net/http"

	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/membership/handlers"
)

func GetHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/membership", "membership", http.MethodGet, handlers.MembershipHandler),
	}
	return &hdrMenus
}