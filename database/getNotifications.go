package database

import (
	"context"
	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

/*GetNotifications function that will read the toots from a profile*/
func GetNotifications(UserID string, page int64) ([]*models.Notification, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("notification")

	var results []*models.Notification
	isStand, _ := IsStandard(UserID)
	var condition bson.M
	if isStand {
		condition = bson.M{
			"userId": UserID,
		}
	} else {
		date, _ := GetDate(UserID)
		condition = bson.M{
			"userId": UserID,
			"date":   bson.M{"$lte": date.DateFrom},
		}
	}
	opti := options.Find()
	opti.SetLimit(5)
	opti.SetSort(bson.D{{Key: "date", Value: -1}})
	opti.SetSkip((page - 1) * 5)

	pointer, err := col.Find(ctx, condition, opti)
	if err != nil {
		return results, false
	}
	for pointer.Next(context.TODO()) {
		var record models.Notification
		err := pointer.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true
}
