package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetCommentsFromToot function that will read the toots from a profile*/
func GetCommentsFromToot(UserID string, TootID string, page int64) ([]*models.ReturnComments, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("comment")

	var results []*models.ReturnComments
	isStand, _ := IsStandard(UserID)
	var condition bson.M
	if isStand {
		condition = bson.M{
			"tootId": TootID,
			"delete": false,
		}
	} else {
		date, _ := GetDate(UserID)
		condition = bson.M{
			"tootId": TootID,
			"delete": false,
			"date":   bson.M{"$lte": date.DateFrom},
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
		var record models.ReturnComments
		err := pointer.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true
}
