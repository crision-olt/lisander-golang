package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertUserReport user to insert in the database the report */
func InsertUserReport(u models.GetUserReport, userID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoCN.Database("lisander")
	col := database.Collection("userReport")
	checkExist, err := CheckReportedUserExistOnReports(u.ReportedUserID)
	if err != nil {
		return false, err
	}
	var r models.Report
	r.UserID = userID
	r.Description = u.Description
	r.Date = time.Now()
	if checkExist {

		m := GetUserReport(u.ReportedUserID)
		m.Reports = append(m.Reports, r)

		condition := bson.M{
			"_id": m.ID,
		}
		updateString := bson.M{
			"$set": m,
		}
		_, err = col.UpdateOne(ctx, condition, updateString)
		if err != nil {
			return false, err
		}
		return true, nil
	} else {
		user, _ := GetUser(u.ReportedUserID)
		user.Password = ""
		var m models.UserReport
		m.IDReportedUser = user.ID.Hex()
		m.ReportedUser = user
		m.Reports = append(m.Reports, r)
		m.Solved = false
		m.Date = time.Now()
		result, err := col.InsertOne(ctx, m)
		if err != nil {
			return false, err
		}

		_, _ = result.InsertedID.(primitive.ObjectID)
		return true, nil
	}

}
