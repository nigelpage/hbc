package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type MatchStore struct {
	Matches []struct {
		Competition struct {
			Name        string `json:"name"`
			BowlslinkID string `json:"bowlslinkId"`
		} `json:"competition"`
		DutySelector struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
		} `json:"dutySelector"`
		Round struct {
			Number   int    `json:"number"`
			PlayedAt string `json:"playedAt"`
			PlayedOn struct {
				Title string `json:"title"`
				Details string `json:"details"`
			} `json:"playedOn"`
			Venue    string `json:"venue"`
			Opponent string `json:"opponent"`
		} `json:"round"`
		Side struct {
			Updated string `json:"updated"`
			Name    string `json:"name"`
			Manager string `json:"manager"`
			Teams []struct {
				Duty string `json:"duty,omitempty"`
				Bowlers []struct {
					Position string `json:"position"`
					Name     string `json:"name"`
					Role     string `json:"role"`
				} `json:"team"`
				Shots struct {
					For     int `json:"for"`
					Against int `json:"against"`
				} `json:"shots"`
			} `json:"teams"`
		} `json:"side"`
	} `json:"matches"`
}

type Teams []struct {
	Duty string `json:"duty,omitempty"`
	Bowlers []struct {
		Position string `json:"position"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	} `json:"team"`
    Shots struct {
        For     int `json:"for"`
        Against int `json:"against"`
    } `json:"shots"`
}

type Team struct {
	Duty string `json:"duty,omitempty"`
	Bowlers []struct {
		Position string `json:"position"`
		Name     string `json:"name"`
		Role     string `json:"role"`
	} `json:"team"`
    Shots struct {
        For     int `json:"for"`
        Against int `json:"against"`
    } `json:"shots"`
}

// type Teams []Team

const (
  timeStoreFormat = "2006-01-02T15:04"
  timeDisplayFormat = "Monday 02 January 2006 03:04PM"
  timeDisplayFormatShort = "02 Jan 03:04PM"
)

func formatBowlslinkURI(bowlslinkId string) string {
	bowlslinkURI :=  fmt.Sprintf("href=https://results.bowlslink.com.au/competition/%s#ladder", bowlslinkId)
	return bowlslinkURI
}

func formatPhoneNumberURI(phoneNumber string) string {
	trimmed := strings.Join(strings.Fields(strings.TrimSpace(phoneNumber)), "")
	tel := fmt.Sprintf("href=tel:+61%s", strings.Replace(trimmed, "0", "", 1))
	return tel
}

func formatTime(timeStr string, shortTimeDisplay bool) string {
	t, err := time.Parse(timeStoreFormat, timeStr)
	if err != nil {
		return "Invalid time format!"
	}
	
	if shortTimeDisplay {
		return t.Format(timeDisplayFormatShort)
	}

	return t.Format(timeDisplayFormat)
}

func hasResults(teams Teams) bool {
	if calculateSidePointsFor(teams) > 0 || calculateSidePointsAgainst(teams) > 0 {
		return true
	}
	
	return false
}

func calculateSidePointsFor(teams Teams) int {
	totalFor := 0
	
	for _, team := range teams {
		totalFor += team.Shots.For
		}

	return totalFor
}

func calculateSidePointsAgainst(teams Teams) int {
	totalAgainst := 0
	
	for _, team := range teams {
		totalAgainst += team.Shots.Against
		}

	return totalAgainst
}

func calculateSidePoints(teams Teams) int {
	teamWins := 0
	teamLosses := 0
	teamDraws := 0
	totalPointsFor := 0
	totalPointsAgainst := 0
	points := 0

	for _, team := range teams {
		pointsFor := team.Shots.For
		pointsAgainst := team.Shots.Against

		totalPointsFor += pointsFor
		totalPointsAgainst += pointsAgainst

		if pointsFor > pointsAgainst {
			teamWins++
		} else if pointsFor < pointsAgainst {
			teamLosses++
		} else {
			teamDraws++
		}
	}

	points = (teamWins * 2) + teamDraws
	if totalPointsFor > totalPointsAgainst {
		points += 10
	} else if totalPointsFor == totalPointsAgainst {
		points +=5
	}

	return points
}

func getSideWinDrawLoss(teams Teams) string {
	totalPointsFor := calculateSidePointsFor(teams)
	totalPointsAgainst := calculateSidePointsAgainst(teams)
	data := ""

	if totalPointsFor > totalPointsAgainst {
		data = "up"
	} else if totalPointsFor < totalPointsAgainst {
		data = "down"
	} else {
		data = "draw"
	}

	return "data-SideResult=" + data
}

func getTeamWinDrawLoss(team Team) string {
data := ""

	if team.Shots.For > team.Shots.Against {
		data = "up"
	} else if team.Shots.For < team.Shots.Against {
		data = "down"
	} else {
		data = "draw"
	}

	return "data-teamResult=" + data
}

func getStoredMatches()	MatchStore {
	// Placeholder function to simulate fetching stored matches
	jsonFile, err := os.Open("Store/W20251108.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return MatchStore{}
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var matchStore MatchStore
	json.Unmarshal(byteValue, &matchStore)

	return matchStore
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	generatedMatches := generateMatches(getStoredMatches())
	baseLayout(generatedMatches).Render(context.Background(), w)
}

func main() {
	/* Setup a handler for static files (e.g. CSS, JS etc...) */
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	/* Setup main handler */
	http.HandleFunc("/", mainHandler)

	/* Start HTTP server */
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}