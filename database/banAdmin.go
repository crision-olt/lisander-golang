package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*BanAdmin validates the report of the user*/
func BanAdmin(ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")

	col := db.Collection("user")

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{
		"_id": objID,
	}
	var user models.User
	col.FindOne(ctx, filter).Decode(&user)
	user.Blocked = true
	updateString := bson.M{
		"$set": user,
	}
	col.UpdateOne(ctx, filter, updateString)
	col = db.Collection("relation")
	filter =
		bson.M{"userId": ID}

	col.DeleteMany(ctx, filter)
	filter =
		bson.M{"userRelationId": ID}
	col.DeleteMany(ctx, filter)
	col = db.Collection("comment")
	filter =
		bson.M{"userId": ID}
	col.DeleteMany(ctx, filter)

	return true, nil
}
