package dbstore

import (
	"github.com/xuri/excelize/v2"
)

func UploadMembers(spreadheet string) (int, int, error) {
	ss, err := excelize.OpenFile(spreadheet)
	if err != nil {
		return 0, 0, err
	}

	defer ss.Close()
	return 0, 0, nil
}