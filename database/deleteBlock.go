package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
)

/*DeleteBlock deletes a relation between users*/
func DeleteBlock(t models.Block) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("block")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
