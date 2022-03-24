package database

import (
	"context"
	"time"
)

/*IsStandard modify database update date*/
func IsStandard(ID string) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dates, err := GetDate(ID)
	if err != nil {
		return false, err
	}
	if dates.UpdateDate == 0 {
		return true, nil
	} else {
		return false, nil
	}
}
