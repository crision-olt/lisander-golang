package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetMyBlocks allows search users by his name*/
func GetMyBlocks(userID string) ([]primitive.ObjectID, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("block")
	condition := bson.M{
		"userId": userID,
	}

	opti := options.Find()
	var userIds []primitive.ObjectID
	cursor, err := col.Find(ctx, condition, opti)
	if err != nil {
		return userIds, false
	}

	for cursor.Next(context.TODO()) {
		var record models.Block
		cursor.Decode(&record)
		var objID primitive.ObjectID
		objID, _ = primitive.ObjectIDFromHex(record.BlockedUserID)

		userIds = append(userIds, objID)
	}

	return userIds, true
}
