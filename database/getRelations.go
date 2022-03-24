package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetRelations allows search users by his name*/
func GetRelations(userID string, typ string) ([]string, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("relation")
	var condition primitive.M
	if typ == "followings" {
		condition = bson.M{
			"userId": userID,
		}
	} else if typ == "followers" {
		condition = bson.M{
			"userRelationId": userID,
		}
	}
	opti := options.Find()
	var userIds []string
	cursor, err := col.Find(ctx, condition, opti)
	if err != nil {
		return userIds, false
	}

	for cursor.Next(context.TODO()) {
		var record models.Relation
		cursor.Decode(&record)
		var objID string
		if typ == "followings" {
			objID = record.UserRelationID
		} else if typ == "followers" {
			objID = record.UserID
		}
		userIds = append(userIds, objID)
	}

	return userIds, true
}
