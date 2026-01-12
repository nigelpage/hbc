package excel

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/xuri/excelize/v2"

	"github.com/nigelpage/hbc/common/store/db"
)

type MemberBase struct {
	CardID int
	FirstName string
	LastName string
}

type UpdateMembersResult struct {
	Added []MemberBase
	Updated []MemberBase
	Deactivated []MemberBase
	Reactivated []MemberBase
}

type member struct {
	card_id int
	first_name string
	last_name string
	is_bowling_member bool
	is_life_member bool
	email_address string
	phone_number string
}

// Extract all the members from the spreadsheet and validate them
func validateMembers(spreadsheet string) ([]member, error) {
ss, err := excelize.OpenFile(spreadsheet)
	if err != nil {
		return nil, fmt.Errorf("Unable to open spreadsheet '%s' %w", spreadsheet, err)
	}

	defer ss.Close()

	membersInSpreadsheet := []member{}

	// Sheets and required columns needed in spreadsheet
	
	sheets := map[string]map[string]int {
		"BOWLING" : {
			"Card ID" : -1,
			"Email Address" : -1,
			"First Name" : -1,
			"Last Name" : -1,
			"SMS Phone Number" : -1,
		},
		"SOCIAL" : {
			"Card ID" : -1,
			"Email Address" : -1,
			"First Name" : -1,
			"Last Name" : -1,
			"SMS Phone Number" : -1,
		},
	}

	// Validate each of the required sheets
	for sheet, fields := range sheets {
	
		// Get all the rows in the sheet
		rows, err := ss.GetRows(sheet)
		if err != nil {
			return nil, fmt.Errorf("Missing sheet %s %w", sheet, err)
		}

		// Check for empty sheet (ignore header row)
		if len(rows) <= 1 {
			return nil, fmt.Errorf("Required sheets cannot be empty")
		}

		// Get header row
		hRow := rows[0];

		// Find the column headers we need
		for index, header := range hRow {
			if _, ok := fields[header]; ok {
				if fields[header] == -1 {
					fields[header] = index
				} else {
					return nil, fmt.Errorf("Duplicate column '%s' in sheet '%s'", header, sheet)
				}
			}
		}

		// Check that all required columns have been found
		for header, index := range fields {
			if index == -1 {
				return nil, fmt.Errorf("Required column '%s' not found in sheet '%s'", header, sheet)
			}
		}
	}

	// Validation successful!
	// All sheets with required columns are present

	// Extract the member details from the spreadsheet

	for sheet, fields := range sheets {
		rows, _ := ss.GetRows(sheet)

		cardIDIndex := fields["Card ID"]
		firstNameIndex := fields["First Name"]
		lastNameIndex := fields["Last Name"]
		emailIndex := fields["Email Address"]
		phoneIndex := fields["SMS Phone Number"]

		// The first empty row indicates that you've hit the end of the rows in use
		for i := 1; rows[i] != nil; i++ {
			cardNo, err := strconv.Atoi(rows[i][cardIDIndex])
			if err != nil {
				return nil, fmt.Errorf("Invalid card number for '%s %s'", rows[i][firstNameIndex], rows[i][lastNameIndex])
			}

			mbr := member {
				card_id: cardNo,
				first_name: rows[i][firstNameIndex],
				last_name: rows[i][lastNameIndex],
				email_address: rows[i][emailIndex],
				phone_number: rows[i][phoneIndex],
			}

			if len(mbr.last_name) > 0 && mbr.last_name[len(mbr.last_name) - 1] == '*' {
				mbr.last_name = strings.TrimSuffix(mbr.last_name, "*")
				mbr.is_life_member = true
			}

			if sheet == "BOWLING" {
				mbr.is_bowling_member = true
			} else {
				mbr.is_bowling_member = false
			}

			// Append member to list of members in spreadsheet
			membersInSpreadsheet = append(membersInSpreadsheet, mbr)
		}
	}

	// All member details extracted from the spreadsheet

	return membersInSpreadsheet, nil
}

func UploadMembers(ctx context.Context, dbPool *pgxpool.Pool, spreadsheet string) (*UpdateMembersResult, error) {

	membersInSpreadsheet, err := validateMembers(spreadsheet)

	// Temporary	
	fmt.Printf("%d members in uploaded spreadsheet", len(membersInSpreadsheet))

	result := UpdateMembersResult{[]MemberBase{}, []MemberBase{}, []MemberBase{}, []MemberBase{}}

	// Now see what needs updating
	q := db.New(dbPool)

	existingMembers, err := q.GetMembers(ctx)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve members from database %w", err)
	}
	fmt.Printf("%v", existingMembers)

	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("Unable to start to transaction for updates %w", err)
	}
	defer tx.Rollback(ctx)

	// Apply updates

	tx.Commit(ctx)

	return &result, nil
}