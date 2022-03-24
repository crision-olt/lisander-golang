package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CountNotifications(UserID string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("notification")

	var results []*models.Notification
	isStand, _ := IsStandard(UserID)
	var condition bson.M
	if isStand {
		condition = bson.M{
			"userId":  UserID,
			"watched": "true",
		}
	} else {
		date, _ := GetDate(UserID)
		condition = bson.M{
			"userId":  UserID,
			"watched": "true",
			"date":    bson.M{"$lte": date.DateFrom},
		}
	}
	opti := options.Find()

	pointer, err := col.Find(ctx, condition, opti)
	if err != nil {
		return 0, err
	}
	for pointer.Next(context.TODO()) {
		var record models.Notification
		err := pointer.Decode(&record)
		if err != nil {
			return 0, err
		}
		results = append(results, &record)
	}
	return int64(len(results)), nil
}
