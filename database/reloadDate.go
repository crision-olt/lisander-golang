package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/crision98/lisander-golang-backend/models"
)

/*ReloadDate modify database update date*/
func ReloadDate(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var err error
	var dates models.User
	db := MongoCN.Database("lisander")
	col := db.Collection("user")
	objID, _ := primitive.ObjectIDFromHex(ID)

	col.FindOne(ctx, bson.M{"_id": bson.M{"$eq": objID}}).Decode(&dates)
	val, _ := CanReloadDate(ID)
	if val {
		if dates.DateFrom.Day() != time.Now().Day() || dates.DateFrom.Month() != time.Now().Month() || dates.DateFrom.Year() != time.Now().Year() {
			dates.UpdateDateToday = dates.UpdateDate
		}
		dates.UpdateDateToday--
	} else {
		return false, nil
	}
	dates.DateFrom = time.Now()
	updateString := bson.M{
		"$set": dates,
	}

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err = col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
