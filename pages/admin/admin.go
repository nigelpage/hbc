package admin

import (
	"net/http"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/admin/handlers"
)

func GetHeaderMenusAndHandlers() *[]common.HeaderMenu {
	hmahs := []common.HeaderMenu {
		{Url: "/admin", Text: "admin", Method: http.MethodGet, Handler: handlers.AdminHandler},
	}
	return &hmahs
}