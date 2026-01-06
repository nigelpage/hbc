package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/pages/pennant/templates"
	store "github.com/nigelpage/hbc/store/json"
)

func getStoredMatches(comp string) store.MatchStore {
	// Placeholder function to simulate fetching stored matches
	pennantCompetitionStore := fmt.Sprintf("./store/json/%s20251108.json", string(comp[0]))
	jsonFile, err := os.Open(pennantCompetitionStore)
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return store.MatchStore{}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var matchStore store.MatchStore
	json.Unmarshal(byteValue, &matchStore)

	return matchStore
}

func createMainPageFromTemplate(competition string) templ.Component {

	store := getStoredMatches(competition)
	generatedMatches := templates.GenerateMatches(store, templates.Icons)
	return templates.BaseLayout(generatedMatches, templates.Icons)
}

/* Main Pennant page handler */
func WeekendCompetitionHandler(ctx echo.Context) error {
	return templateRenderer(ctx, http.StatusOK, createMainPageFromTemplate("Weekend"))
}

func CompetitionHandler(ctx echo.Context) error {
	comp := ctx.Param("competition")
	return templateRenderer(ctx, http.StatusOK, createMainPageFromTemplate(comp))
}