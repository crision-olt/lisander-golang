package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*DeleteNotification deletes a notification*/
func DeleteNotification(opt primitive.M) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("notification")

	_, err := col.DeleteOne(ctx, opt)
	if err != nil {
		return false, err
	}
	return true, nil
}
