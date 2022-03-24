package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*GetUserReport looking for profile in DB*/
func GetUserReport(ID string) *models.UserReport {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("userReport")

	condition := bson.M{
		"idReportedUser": ID,
		"solved":         false,
	}
	var report *models.UserReport
	col.FindOne(ctx, condition).Decode(&report)

	return report

}
