package database

import (
	"context"
	"time"
)

/*IsAdmin modify database update date*/
func IsAdmin(Admin int) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if Admin > 0 {
		return true, nil
	}
	return false, nil
}
