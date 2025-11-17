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
		Competition  string `json:"competition"`
		DutySelector struct {
			Name  string `json:"name"`
			Phone string `json:"phone"`
		} `json:"dutySelector"`
		Round struct {
			Number   int    `json:"number"`
			PlayedAt string `json:"playedAt"`
			Venue    string `json:"venue"`
			Opponent string `json:"opponent"`
		} `json:"round"`
		Side struct {
			Updated string `json:"updated"`
			Name    string `json:"name"`
			Manager string `json:"manager"`
			Teams   []struct {
				Duty string `json:"duty"`
				Bowlers []struct {
					Position string `json:"position"`
					Name     string `json:"name"`
					Role     string `json:"role"`
				} `json:"team"`
			} `json:"teams"`
		} `json:"side"`
	} `json:"matches"`
}

const (
  timeStoreFormat = "2006-01-02T15:04"
  timeDisplayFormat = "Monday 02 January 2006 03:04PM"
  timeDisplayFormatShort = "02 Jan 03:04PM"
)

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