package membership

import (
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/pages/header"
	"github.com/nigelpage/hbc/pages/membership/handlers"
)

func getHeaderMenus() *[]header.HeaderMenu {
	hdrMenus := []header.HeaderMenu{
		*header.NewHeaderMenu("/membership", "membership", http.MethodGet, handlers.MembershipHandler, nil),
	}
	return &hdrMenus
}

func NewMembershipPage(dbPool *pgxpool.Pool) (*common.PageDetails, error) {
	return common.NewPage("Membership", "Membership information and application", "v0.0.1-alpha",
						  dbPool, getHeaderMenus(), nil)
}