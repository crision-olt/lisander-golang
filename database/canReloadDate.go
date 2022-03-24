package database

import (
	"context"
	"time"
)

/*CanReloadDate modify database update date*/
func CanReloadDate(ID string) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	dates, err := GetDate(ID)
	if err != nil {
		return false, err
	}
	if dates.DateFrom.Day() == time.Now().Day() && dates.DateFrom.Month() == time.Now().Month() && dates.DateFrom.Year() == time.Now().Year() {
		if dates.UpdateDateToday == 0 {
			return false, nil
		}
		return true, nil
	} else {
		return true, nil
	}
}
