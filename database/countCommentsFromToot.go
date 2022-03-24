package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

/*CountCommentsFromToot looking for profile in DB*/
func CountCommentsFromToot(UserID string, TootID string) (int64, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("comment")
	date, _ := GetDate(UserID)
	condition := bson.M{
		"tootId": TootID,
		"date":   bson.M{"$lt": date.DateFrom},
		"hide":   false,
		"delete": false,
	}
	count, err := col.CountDocuments(ctx, condition)

	if err != nil {
		return count, err
	}
	return count, nil

}
