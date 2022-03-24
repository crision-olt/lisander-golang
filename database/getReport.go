package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*GetReport looking for profile in DB*/
func GetReport(ID string) (models.UserReport, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("userReport")

	var profile models.UserReport
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"solved": false,
	}
	err := col.FindOne(ctx, condition).Decode(&profile)
	if err != nil {
		return profile, err
	}
	return profile, nil

}
