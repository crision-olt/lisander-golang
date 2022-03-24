package database

import (
	"context"
	"time"
)

/*IsSuperAdmin modify database update date*/
func IsSuperAdmin(Admin int) (bool, error) {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if Admin > 1 {
		return true, nil
	}
	return false, nil
}
