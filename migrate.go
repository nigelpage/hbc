package main

// import (
// 	"fmt"
// 	"os"
// 	"io"
// 	"encoding/json"

// 	"github.com/jackc/pgx/v5/pgxpool"

// 	store "github.com/nigelpage/hbc/store/json"
// 	dbstore "github.com/nigelpage/hbc/store/db"
// )

// func migrateFromJsonToDB(pool *pgxpool.Pool, queries *dbstore.Queries) error {
// 	// Load JSON file
// 	jsonFile, err := os.Open("pages/pennant/store/%W20251108.json")
// 	if err != nil {
// 		return fmt.Errorf("Error opening JSON file: %w", err)
// 	}
// 	defer jsonFile.Close()

// 	byteValue, _ := io.ReadAll(jsonFile)

// 	var matchStore store.MatchStore
// 	json.Unmarshal(byteValue, &matchStore)

// 	//JSON file now loaded!
// 	// Migrate data to DB

// for _, match := range matchStore.Matches {
// 	_, err := queries.CreateMatch(context.Background(), dbstore.CreateMatchParams{
// 		Competition: match.Competition,
// 		HomeTeam:    match.HomeTeam,
// 		AwayTeam:    match.AwayTeam,
// 		Date:        match.Date,
// 		HomeScore:   match.HomeScore,
// 		AwayScore:   match.AwayScore,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("Error creating match in DB: %w", err)
// 	}
// }

// 	return nil
// }
