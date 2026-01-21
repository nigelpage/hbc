package index

import (
	"net/http"

	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/index/handlers"
)

func GetHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/", "home", http.MethodGet, handlers.IndexHandler, nil),
	}
	return &hdrMenus
}
