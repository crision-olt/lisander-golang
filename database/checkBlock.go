package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckBlock check the block between two user*/
func CheckBlock(t models.Block) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("block")

	condition := bson.M{
		"userId":        t.UserID,
		"blockedUserId": t.BlockedUserID,
	}

	var result models.Block
	err := col.FindOne(ctx, condition).Decode(&result)
	if result != (models.Block{}) {
		return true, err
	}
	condition = bson.M{
		"userId":        t.BlockedUserID,
		"blockedUserId": t.UserID,
	}
	err = col.FindOne(ctx, condition).Decode(&result)
	if result != (models.Block{}) {
		return true, err
	}
	return false, nil
}
