package admin

import (
	"net/http"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/admin/handlers"
)

func GetHeaderMenus() *[]common.HeaderMenu {
	hdrMenus := []common.HeaderMenu{
		*common.NewHeaderMenu("/admin", "admin", http.MethodGet, handlers.AdminHandler),
	}
	return &hdrMenus
}