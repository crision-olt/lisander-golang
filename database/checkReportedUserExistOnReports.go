package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*CheckReportedUserExistOnReports looking for profile in DB*/
func CheckReportedUserExistOnReports(ID string) (bool, error) {

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
	if report != nil {
		return true, nil
	}
	return false, nil

}
