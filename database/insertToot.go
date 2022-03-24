package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertToot records toot in DB*/
func InsertToot(t models.InsertToot) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("toot")
	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return string(""), false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
