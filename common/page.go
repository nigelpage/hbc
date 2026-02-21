package common

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/nigelpage/hbc/pages/header"
	"golang.org/x/mod/semver"
)

type PageDetails struct {
	Title       string
	Description string
	Version	 string
	dbPool *pgxpool.Pool
	headerMenus *[]header.HeaderMenu
	adminPageHandler func(ctx echo.Context) error
}

func NewPage(title, description, version string,
			 dbPool *pgxpool.Pool,
			 headerMenus *[]header.HeaderMenu,
			 adminPageHandler func(ctx echo.Context) error) (*PageDetails, error) {
	if (!semver.IsValid(version)) {
		return nil, fmt.Errorf("Invalid version format for %s page: %s", title, version)
	}

	return &PageDetails{
		Title:       title,
		Description: description,
		Version:     version,
		dbPool:      dbPool,
		headerMenus: headerMenus,
		adminPageHandler: adminPageHandler,
	}, nil
}

func (pd *PageDetails) DbPool() *pgxpool.Pool {
	return pd.dbPool
}
func (pd *PageDetails) HeaderMenus() *[]header.HeaderMenu {
	return pd.headerMenus
}

func (pd *PageDetails) AdminPageHandler() func(ctx echo.Context) error {
	return pd.adminPageHandler
}