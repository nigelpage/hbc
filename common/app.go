package common

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
)

type App struct {
	Echo	*echo.Echo
}

// Create a new Echo server application
func NewApp(echo *echo.Echo) (*App) {
	// Initialise Echo web server
	app := &App {
		Echo:	echo,
	}

	return app
}

// Start the Echo server application
func (app *App) Start(address string) error {
	quit, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	s := http.Server{
		Addr:    address,
		Handler: app.Echo,
		//ReadTimeout: 30 * time.Second, // customize http.Server timeouts
	}
	go func(srv *http.Server) {
		if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			app.Echo.Logger.Fatal(err)
		}
		cancel() // in case server returns before ctrl+c
	}(&s)

	app.Echo.Logger.Infof("Started server on '%s'", address)
	
	// Wait until interrupt signal to start shutdown
	<-quit.Done()

	// start gracefully shutdown with a timeout of 10 seconds.
	ctx, cancelGC := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelGC()

	if err := s.Shutdown(ctx); err != nil {
		app.Echo.Logger.Fatal(err)
	}
	app.Echo.Logger.Info("Shutting down server")

	return nil
}