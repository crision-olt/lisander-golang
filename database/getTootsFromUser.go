package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetTootsFromUser function that will read the toots from a profile*/
func GetTootsFromUser(UserID string, ID string, page int64) ([]*models.ReturnToots, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("toot")

	var results []*models.ReturnToots
	isStand, _ := IsStandard(UserID)
	var condition bson.M
	if isStand || UserID == ID {
		condition = bson.M{
			"userId":  ID,
			"hide":    false,
			"deleted": false,
		}
	} else {
		date, _ := GetDate(UserID)
		condition = bson.M{
			"userId":  ID,
			"hide":    false,
			"deleted": false,
			"date":    bson.M{"$lte": date.DateFrom},
		}
	}

	opti := options.Find()
	opti.SetLimit(20)
	opti.SetSort(bson.D{{Key: "date", Value: -1}})
	opti.SetSkip((page - 1) * 20)

	pointer, err := col.Find(ctx, condition, opti)
	if err != nil {
		return results, false
	}
	for pointer.Next(context.TODO()) {
		var record models.ReturnToots
		err := pointer.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true
}
