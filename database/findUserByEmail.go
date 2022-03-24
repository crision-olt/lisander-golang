package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*FindUserByEmail allows search users by his name*/
func FindUserByEmail(userID string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	var results []*models.User
	db := MongoCN.Database("lisander")
	col := db.Collection("user")
	var ids []primitive.ObjectID
	ids, _ = GetMyBlocks(userID)
	objId, _ := primitive.ObjectIDFromHex(userID)
	ids = append(ids, objId)
	cursor, err := col.Find(ctx, bson.M{
		"_id":     bson.M{"$nin": ids},
		"blocked": false,
	})

	if err != nil {
		return results, false
	}
	for cursor.Next(ctx) {
		var s models.User
		cursor.Decode(&s)
		results = append(results, &s)
	}
	return results, true
}
