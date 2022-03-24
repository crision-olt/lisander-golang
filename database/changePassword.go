package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/crision98/lisander-golang-backend/models"
)

/*ChangePassword allows to modify the user password*/
func ChangePassword(u models.ChangePassword, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("lisander")
	col := db.Collection("user")

	record := make(map[string]interface{})
	record["password"], _ = EncryptPassword(u.NewPassword)
	updateString := bson.M{
		"$set": record,
	}
	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)
	if err != nil {
		return false, err
	}

	return true, nil
}
