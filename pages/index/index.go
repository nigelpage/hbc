package index

import (
	"net/http"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/index/handlers"
)

func GetHeaderMenus() *[]common.HeaderMenu {
	hdrMenus := []common.HeaderMenu{
		*common.NewHeaderMenu("/", "home", http.MethodGet, handlers.IndexHandler),
	}
	return &hdrMenus
}
