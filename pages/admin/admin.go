package admin

import (
	"net/http"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/admin/handlers"
)

func GetHeaderMenusAndHandlers() *[]common.HeaderMenuAndHandler {
	hmahs := []common.HeaderMenuAndHandler{
		{Url: "/admin", Text: "admin", Method: http.MethodGet, Handler: handlers.AdminHandler},
	}
	return &hmahs
}