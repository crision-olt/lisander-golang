package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetReports function that will read the toots from a profile*/
func GetReports(page int64) ([]*models.UserReport, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("userReport")

	var results []*models.UserReport

	condition := bson.M{"solved": false}

	opti := options.Find()
	opti.SetLimit(20)
	opti.SetSort(bson.D{{Key: "date", Value: -1}})
	opti.SetSkip((page - 1) * 20)

	pointer, err := col.Find(ctx, condition, opti)
	if err != nil {
		return results, false
	}
	for pointer.Next(context.TODO()) {
		var record models.UserReport
		err := pointer.Decode(&record)
		if err != nil {
			return results, false
		}
		results = append(results, &record)
	}
	return results, true
}
