package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*GetToot looking for profile in DB*/
func GetToot(ID string) (models.ReturnToots, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("toot")

	var toot models.ReturnToots
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}
	err := col.FindOne(ctx, condition).Decode(&toot)
	if err != nil {
		return toot, err
	}
	return toot, nil

}
