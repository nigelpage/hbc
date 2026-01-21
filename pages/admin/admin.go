package admin

import (
	"net/http"

	"github.com/nigelpage/hbc/pages/admin/handlers"
	"github.com/nigelpage/hbc/pages/header"
)

func GetHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/admin", "admin", http.MethodGet, handlers.AdminHandler),
	}
	return &hdrMenus
}