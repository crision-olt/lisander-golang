package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertComment records comment in DB*/
func InsertComment(t models.InsertComment) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("comment")

	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return string(""), false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	col = db.Collection("toot")
	var toot models.InsertToot
	objID2, _ := primitive.ObjectIDFromHex(t.TootID)
	col.FindOne(ctx, bson.M{
		"_id": objID2,
	}).Decode(&toot)
	if toot.UserID != t.UserID {
		CreateNotification(models.Notification{TypeID: 2, UserID: toot.UserID, FromUserID: t.UserID, FromTootID: t.TootID, FromCommentID: objID.Hex()})
	}
	return objID.String(), true, nil
}
