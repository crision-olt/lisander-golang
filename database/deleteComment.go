package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*DeleteComment deletes a toot*/
func DeleteComment(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("comment")

	objID, _ := primitive.ObjectIDFromHex(ID)
	condition := bson.M{
		"_id":    objID,
		"userId": UserID,
	}
	var comment *models.InsertComment
	col.FindOne(ctx, condition).Decode(&comment)
	comment.Delete = true
	updateString := bson.M{
		"$set": comment,
	}
	_, err := col.UpdateOne(ctx, condition, updateString)
	DeleteNotification(bson.M{
		"fromCommentId": ID,
	})
	return err
}
