package main

import (
	"fmt"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/pennant/handlers"
)

type Password struct {
	Hash string `json:"password"`
}

func checkHash(password Password) bool {
	fmt.Printf("Password hash: %s/n", password.Hash)
	return true;
}

// func authenticationHandler(ctx echo.Context) error {
// // Source - https://stackoverflow.com/a
// // Posted by transistor
// // Retrieved 2025-12-13, License - CC BY-SA 3.0

// // Read the Body content
// 	var bodyBytes []byte
// 	if ctx.Request().Body != nil {
//     	bodyBytes, _ = io.ReadAll(ctx.Request().Body)
// 	}

// // Restore the io.ReadCloser to its original state
// 	ctx.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

// // Continue to use the Body, like Binding it to a struct:
// 	password := new(Password)
// 	error := ctx.Bind(password)
// 	if error != nil {
// 		if checkHash(*password) {
// 			return nil
// 		}
// 	}

// 	ctx.Response().Writer.Write([]byte(LockedEdit(icons, "Invalid password!")))
// }

func main() {
	app := echo.New()
	/* Setup a handler for static files (e.g. CSS, JS etc...) */
	app.Static("/static", "./assets")
	
	/* Setup main handler */
	app.GET("/", handlers.WeekendCompetitionHandler)
	app.GET("/:competition", handlers.CompetitionHandler)
	// app.POST("/authenticate", authenticationHandler)

	/* Start HTTP server */
	app.Logger.Fatal(app.Start(":4000"))
}