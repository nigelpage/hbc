package pennant

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/pennant/handlers"
)

func getHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/pennant/:competition", "pennant", http.MethodGet, handlers.PennantHandler,
							  map[string]string{"competition": "weekend"}),
		*header.NewHeaderMenu("/pennant", "", http.MethodGet, handlers.PennantHandler, nil),
	}
	return &hdrMenus
}

func NewPennantPage(dbPool *pgxpool.Pool) (*common.PageDetails, error) {
	return common.NewPage("Pennant", "Pennant sides selections and results", "v0.0.1-alpha",
						  dbPool, getHeaderMenus(), nil)
}