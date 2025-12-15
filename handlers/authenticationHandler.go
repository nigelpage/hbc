package handlers

import (
	"fmt"
	"net/http"

	// "github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/nigelpage/pennant/templates"
	"golang.org/x/crypto/bcrypt"
)

var hash = "$2a$14$ab0Q4VcHrW.4ZSWWbejZQeu.1KRjfBQyxSF68TfCTSkCStsNmZ7OO" // gr33nW*y3135

// func HashPassword(password string) (string, error) {
//     bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//     return string(bytes), err
// }

func AuthenticationHandler(ctx echo.Context) error {
	pwd := ctx.Request().PostFormValue("password")
	if pwd == "" {
		return fmt.Errorf("Empty password not allowed!")
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		return fmt.Errorf("Invalid password!")
	}

	return templateRenderer(ctx, http.StatusOK, templates.UnlockedEdit(templates.Icons))
}

func LockAuthenticationHandler(ctx echo.Context) error {
	return templateRenderer(ctx, http.StatusOK, templates.LockedEdit(templates.Icons, ""))
}