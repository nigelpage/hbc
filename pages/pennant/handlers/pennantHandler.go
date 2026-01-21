package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	js "github.com/nigelpage/hbc/common/store/json"
	ht "github.com/nigelpage/hbc/pages/header/templates"
	"github.com/nigelpage/hbc/pages/pennant/templates"
)

func getStoredMatches(comp string) js.MatchStore {
	// Placeholder function to simulate fetching stored matches
	pennantCompetitionStore := fmt.Sprintf("./common/store/json/%s20251108.json", string(comp[0]))
	jsonFile, err := os.Open(pennantCompetitionStore)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return js.MatchStore{}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var matchStore js.MatchStore
	json.Unmarshal(byteValue, &matchStore)

	return matchStore
}

/* Main Pennant page handler */

func PennantHandler(ctx echo.Context) error {
	comp := ctx.Param("competition")
	if (comp == "") {
		comp = "weekend"
	} else {
		comp = strings.ToLower(comp)
	}
	store := getStoredMatches(comp)

	return common.TemplateRenderer(ctx, http.StatusOK, ht.CreatePageFromTemplate(ctx, templates.PennantLayout(store, templates.Icons)))
}