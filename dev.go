package main

import (
	"fmt"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/nigelpage/hbc/common"
	"github.com/nigelpage/hbc/common/store/excel"
)
	// Migrate from Json to database
	// err = migrateFromJsonToDB(app.Pool, app.Queries)

	// Upload members from Excel spreadsheet
	func whileDev(ctx context.Context, dbPool *pgxpool.Pool) {
	results, err := excel.UploadMembers(ctx, dbPool, "./common/store/excel/Members Draw list 09.01.2026.xlsx")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d member%s added, %d member%s updated, %d member%s deactivated\n",
										results.Added, common.Pluralise(results.Added),
										results.Updated, common.Pluralise(results.Updated),
										results.Deactivated, common.Pluralise(results.Deactivated))

	// Load ticker messages
	
}

