package pennant

import (
	"net/http"

	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/pennant/handlers"
)

func GetHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/pennant/:competition", "pennant", http.MethodGet, handlers.PennantHandler,
							  map[string]string{"competition": "weekend"}),
		*header.NewHeaderMenu("/pennant", "", http.MethodGet, handlers.PennantHandler, nil),
	}
	return &hdrMenus
}