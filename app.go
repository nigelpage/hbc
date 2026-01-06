package main

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nigelpage/hbc/store/db"

	"github.com/labstack/echo/v4"
)

type App struct {
	Echo	*echo.Echo
	Pool	*pgxpool.Pool
	Queries	*dbstore.Queries
}

func NewApp(echo *echo.Echo, pool *pgxpool.Pool, queries *dbstore.Queries) (*App) {
	// Initialise Echo web server
	app := &App {
		Echo:	echo,
		Pool:   pool,
		Queries: queries,
	}

	return app
}