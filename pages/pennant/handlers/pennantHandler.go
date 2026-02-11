package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/nigelpage/hbc/common"
	js "github.com/nigelpage/hbc/common/store/json"
	ht "github.com/nigelpage/hbc/pages/header/templates"
	"github.com/nigelpage/hbc/pages/pennant/templates"
)

func getStoredMatches(comp string, we time.Time) (js.MatchStore, error) {
	// Placeholder function to simulate fetching stored matches
	pennantCompetitionStore := fmt.Sprintf("./common/store/json/%s%s.json", string(comp[0]), we.Format("20060102"))
	jsonFile, err := os.Open(pennantCompetitionStore)
	if err != nil {
		return js.MatchStore{},
			   fmt.Errorf("No data found for %s competition, week ending %s: %w", comp, we.Format("02-Jan-2006"), err)
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var matchStore js.MatchStore
	json.Unmarshal(byteValue, &matchStore)

	return matchStore, nil
}

/* Main Pennant page handler */

func PennantHandler(ctx echo.Context) error {
	comp := ctx.Param("competition")
	if (comp != "") {
		comp = strings.ToLower(comp)
	} else {
		comp = "weekend"
	}

	layout := "02-Jan-2006"
	we := time.Now()
	we = we.AddDate(0, 0, -int(we.Weekday())+6) // Get the upcoming Saturday

	weekEnding := ctx.QueryParam("weekEnding")

	var err error
	
	if (weekEnding != "") {
		we, err = time.Parse(layout, weekEnding)
		if (err != nil) {
			return echo.NewHTTPError(http.StatusBadRequest,
									 fmt.Sprintf("Invalid weekEnding format: %s - expected format is DD-MMM-YYYY", weekEnding))
		}
	}

	store, err := getStoredMatches(comp, we)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return common.TemplateRenderer(ctx, http.StatusOK, ht.CreatePageFromTemplate(ctx, templates.PennantLayout(store, templates.Icons)))
}