package handlers

import (
	"fmt"
	"io"
	"encoding/json"
	"os"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"

	"github.com/nigelpage/pennant/store"
	"github.com/nigelpage/pennant/templates"
)

func getStoredMatches(comp string) store.MatchStore {
	// Placeholder function to simulate fetching stored matches
	pennantCompetitionStore := fmt.Sprintf("Store/%s20251108.json", string(comp[0]))
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
	return mainPageRenderer(ctx, http.StatusOK, createMainPageFromTemplate("Weekend"))
}

func CompetitionHandler(ctx echo.Context) error {
	comp := ctx.Param("competition")
	return mainPageRenderer(ctx, http.StatusOK, createMainPageFromTemplate(comp))
}

/* Main page renderer */
func mainPageRenderer(ctx echo.Context, statusCode int, cmp templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	
	if err := cmp.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}