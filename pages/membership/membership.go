package membership

import (
	"net/http"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/membership/handlers"
)

func GetHeaderMenus() *[]common.HeaderMenu {
	hdrMenus := []common.HeaderMenu{
		*common.NewHeaderMenu("/membership", "membership", http.MethodGet, handlers.MembershipHandler),
	}
	return &hdrMenus
}