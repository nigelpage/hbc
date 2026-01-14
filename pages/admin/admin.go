package admin

import (
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/admin/handlers"
)

func GetHandlers() []*common.Handler {
	return []*common.Handler{
		common.NewHandler("/admin", "GET", handlers.AdminHandler),
	}
}