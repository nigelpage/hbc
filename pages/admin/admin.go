package admin

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/admin/handlers"
	"github.com/nigelpage/hbc/pages/header"
)

func getHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/admin", "admin", http.MethodGet, handlers.AdminHandler, nil),
	}
	return &hdrMenus
}

func NewAdminPage(dbPool *pgxpool.Pool) (*common.PageDetails, error) {
	return common.NewPage("Admin", "Admin dashboard and management tools", "v0.0.1-alpha",
						  dbPool, getHeaderMenus(), nil)
}