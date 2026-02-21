package index

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/index/handlers"
)

func getHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/", "home", http.MethodGet, handlers.IndexHandler, nil),
	}
	return &hdrMenus
}

func NewIndexPage(dbPool *pgxpool.Pool) (*common.PageDetails, error) {
	return common.NewPage("Home", "Index page", "v0.0.1-alpha",
						  dbPool, getHeaderMenus(), nil)
}
