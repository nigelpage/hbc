package excel

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

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
	Added int
	Updated int
}

type member struct {
	card_id int
	first_name string
	last_name string
	email_address string
	phone_number string
	is_bowling_member bool
	is_life_member bool
	is_financial bool
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
			"Not paid/Part paid" : -1,
		},
		"SOCIAL" : {
			"Card ID" : -1,
			"Email Address" : -1,
			"First Name" : -1,
			"Last Name" : -1,
			"SMS Phone Number" : -1,
			"Not paid/Part paid" : -1,
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
		isFinancialIndex := fields["Part Paid/Not Paid"]

		// The first empty row indicates that you've hit the end of the rows in use
		for i := 1; rows[i] != nil; i++ {
			cardNo, err := strconv.Atoi(rows[i][cardIDIndex])
			if err != nil {
				return nil, fmt.Errorf("Invalid card number for '%s %s'", rows[i][firstNameIndex], rows[i][lastNameIndex])
			}

			// Save the member info required
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
			} else {
				mbr.is_life_member = false
			}

			if sheet == "BOWLING" {
				mbr.is_bowling_member = true
			} else {
				mbr.is_bowling_member = false
			}

			if strings.EqualFold(strings.TrimSpace(rows[i][isFinancialIndex]), "Not Paid") {
				mbr.is_financial = true
			} else {
				mbr.is_financial = false
			}

			// Append member to list of members in spreadsheet
			membersInSpreadsheet = append(membersInSpreadsheet, mbr)
		}
	}

	// All member details extracted from the spreadsheet

	return membersInSpreadsheet, nil
}

func convertSSMemberToDBMember(ssmbr member, at time.Time, isNew bool) *db.Member {
	mbr := db.Member {
			MembershipNumber: int32(ssmbr.card_id),
			FirstName: ssmbr.first_name,
			LastName: ssmbr.last_name,
			Email: db.ConvertStringToPGText(ssmbr.email_address),
			Phone: db.ConvertStringToPGText(ssmbr.phone_number),
			IsBowlingMember: db.ConvertBoolToPGBool(ssmbr.is_bowling_member),
			IsLifeMember: db.ConvertBoolToPGBool(ssmbr.is_life_member),
			IsFinancial: db.ConvertBoolToPGBool(ssmbr.is_financial),
			UpdatedAt: db.ConvertTimeToPGTimestamptz(at),
		}
	if isNew {
		mbr.CreatedAt = db.ConvertTimeToPGTimestamptz(at)
	}
	return &mbr
}

func convertMemberToCreateMemberParams(mbr db.Member) *db.CreateMemberParams {
	return &db.CreateMemberParams {
		MembershipNumber: mbr.MembershipNumber,
		FirstName: mbr.FirstName,
		LastName: mbr.LastName,
		Email: mbr.Email,
		Phone: mbr.Phone,
		IsBowlingMember: mbr.IsBowlingMember,
		IsLifeMember: mbr.IsLifeMember,
		IsFinancial: mbr.IsFinancial,
		IsActive: mbr.IsActive,
	}
}

func convertMemberToUpdateMemberDetailsParams(mbr db.Member) *db.UpdateMemberDetailsParams {
	return &db.UpdateMemberDetailsParams {
		MembershipNumber: mbr.MembershipNumber,
		FirstName: mbr.FirstName,
		LastName: mbr.LastName,
		Email: mbr.Email,
		Phone: mbr.Phone,
		IsBowlingMember: mbr.IsBowlingMember,
		IsLifeMember: mbr.IsLifeMember,
		IsFinancial: mbr.IsFinancial,
		IsActive: mbr.IsActive,
	}
}

func UploadMembers(ctx context.Context, dbPool *pgxpool.Pool, spreadsheet string) (*UpdateMembersResult, error) {

	membersInSpreadsheet, err := validateMembers(spreadsheet)

	// Temporary	
	fmt.Printf("%d members in uploaded spreadsheet\n", len(membersInSpreadsheet))

	// Now see what needs updating
	q := db.New(dbPool)

	existingMembers, err := q.FindMembers(ctx)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve members from database %w", err)
	}

	existingMemberCount := 0

	if existingMembers != nil {
		existingMemberCount = len(existingMembers)
	}
	fmt.Printf("%d existing members in database\n", existingMemberCount)

	var added []db.Member
	var updated []db.Member

	// Ensure all records have the same timestamp
	// Makes identifying an upload run much easier!
	now := time.Now()

	// Check for initial load of members
	if existingMemberCount == 0 {
		for _, uploadedMbr := range membersInSpreadsheet {
			mbrFromUpload := convertSSMemberToDBMember(uploadedMbr, now, true)
			added = append(added, *mbrFromUpload)
		}
	} else {
		// Check for members that need to be added or updated
		for _, uploadedMbr := range membersInSpreadsheet {
		checkIfInDatabaseLoop:
			for _, mbrInDatabase := range existingMembers {
				if uploadedMbr.card_id == int(mbrInDatabase.MembershipNumber) {
					mbrFromUpload := convertSSMemberToDBMember(uploadedMbr, now, false)
					if !mbrInDatabase.Equal(*mbrFromUpload) {
							// Member needs to be updated!
							updated = append(updated, *mbrFromUpload)
							break checkIfInDatabaseLoop
					}
					break checkIfInDatabaseLoop
				}
			}
			// New member needs to be added!
			newMember := convertSSMemberToDBMember(uploadedMbr, now, true)		
			added = append(added, *newMember)
		}
	}

	// Make sure that everything's updated or nothing by wrapping it all in a transaction
	tx, err := dbPool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("Unable to start to transaction for updates %w", err)
	}
	defer tx.Rollback(ctx)

	q = db.New(dbPool).WithTx(tx)

	// Add new members
	if len(added) > 0 {
		for _, mbr := range added {
			_, err := q.CreateMember(ctx, *convertMemberToCreateMemberParams(mbr))
			if err != nil {
				return nil, fmt.Errorf("Failed to add new member '%s %s' %w", mbr.FirstName, mbr.LastName, err)
			}
		}
	}

	// Update members with changed details

	if len(updated) > 0 {
		for _, mbr := range updated {
			err := q.UpdateMemberDetails(ctx, *convertMemberToUpdateMemberDetailsParams(mbr))
			if err != nil {
				return nil, fmt.Errorf("Failed to update member '%s %s' %w", mbr.FirstName, mbr.LastName, err)
			}
		}
	}

	tx.Commit(ctx)

	return &UpdateMembersResult { Added: len(added), Updated: len(updated)}, nil
}