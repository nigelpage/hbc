package db

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func ConvertStringToPGText (str string) pgtype.Text {
	if str == "" {
		return pgtype.Text { String: "", Valid: false }
	}
	return pgtype.Text{ String: str, Valid: true }
}

func ConvertBoolToPGBool (tf bool) pgtype.Bool {
	return pgtype.Bool{ Bool: tf, Valid: true }
}

func ConvertTimeToPGTimestamptz (datetime time.Time) pgtype.Timestamptz {
	return pgtype.Timestamptz{ Time: datetime.UTC(), InfinityModifier: pgtype.Finite , Valid: true }
}

func (mbr1 Member) Equal(mbr2 Member) bool {
	return mbr1.MembershipNumber == mbr2.MembershipNumber &&
		   mbr1.FirstName == mbr2.FirstName &&
		   mbr1.LastName == mbr2.LastName &&
		   mbr1.Email == mbr2.Email &&
		   mbr1.Phone == mbr2.Phone &&
		   mbr1.IsBowlingMember == mbr2.IsBowlingMember &&
		   mbr1.IsLifeMember == mbr2.IsLifeMember &&
		   mbr1.IsFinancial == mbr2.IsFinancial &&
		   mbr1.IsActive == mbr2.IsActive
}