package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetTootsFromRelations read the toots from my followers*/
func GetTootsFromRelations(userID string, page int64) ([]models.ReturnTootsFollowers, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []models.ReturnTootsFollowers
	userIds, _ := GetRelations(userID, "followings")
	db := MongoCN.Database("lisander")
	col := db.Collection("toot")
	isStand, _ := IsStandard(userID)
	var condition bson.M
	if isStand {
		condition = bson.M{
			"userId":  bson.M{"$in": userIds},
			"hide":    false,
			"deleted": false,
		}
	} else {
		date, _ := GetDate(userID)
		condition = bson.M{
			"userId":  bson.M{"$in": userIds},
			"hide":    false,
			"deleted": false,
			"date":    bson.M{"$lte": date.DateFrom},
		}
	}
	opti := options.Find()
	opti.SetLimit(20)
	opti.SetSort(bson.D{{Key: "date", Value: -1}})
	opti.SetSkip((page - 1) * 20)
	cursor, err := col.Find(ctx, condition, opti)

	if err != nil {
		return results, true
	}
	for cursor.Next(ctx) {
		var record models.ReturnTootsFollowers
		cursor.Decode(&record)
		results = append(results, record)
	}
	return results, true
}
