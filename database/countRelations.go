package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*CountRelations allows search users by his name*/
func CountRelations(userID string, typ string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("relation")

	var condition primitive.M
	date, _ := GetDate(userID)
	if typ == "followings" {
		condition = bson.M{
			"userId": userID,
			"date":   bson.M{"$lt": date.DateFrom},
		}
	} else if typ == "followers" {
		condition = bson.M{
			"userRelationId": userID,
			"date":           bson.M{"$lt": date.DateFrom},
		}
	}
	var results []models.Relation
	cursor, _ := col.Find(ctx, condition)

	for cursor.Next(ctx) {
		var s models.Relation
		cursor.Decode(&s)
		results = append(results, s)
	}

	return int64(len(results)), nil
}
