package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteToot deletes a toot*/
func DeleteToot(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("toot")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userId": UserID,
	}
	var toot *models.InsertToot
	col.FindOne(ctx, condition).Decode(&toot)
	toot.Deleted = true
	updateString := bson.M{
		"$set": toot,
	}
	DeleteNotification(bson.M{
		"fromTootId": ID,
	})
	_, err := col.UpdateOne(ctx, condition, updateString)
	return err
}
