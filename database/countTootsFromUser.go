package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*CountTootsFromUser function that will read the toots from a profile*/
func CountTootsFromUser(ID string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("toot")
	date, _ := GetDate(ID)
	condition := bson.M{
		"userId":  ID,
		"date":    bson.M{"$lt": date.DateFrom},
		"hide":    false,
		"deleted": false,
	}

	count, err := col.CountDocuments(ctx, condition)
	if err != nil {
		return count, err
	}

	return count, nil
}
