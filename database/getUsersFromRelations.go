package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetUsersFromRelations search users by relations*/
func GetUsersFromRelations(userID string, page int64, typ string) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var results []*models.User
	userIds, _ := GetRelations(userID, typ)
	var userIdsObject []primitive.ObjectID
	for _, id := range userIds {
		primitive, _ := primitive.ObjectIDFromHex(id)
		userIdsObject = append(userIdsObject, primitive)
	}
	db := MongoCN.Database("lisander")
	col := db.Collection("user")
	condition := bson.M{
		"_id": bson.M{
			"$in": userIdsObject,
		},
	}
	opti := options.Find()
	opti.SetLimit(20)
	opti.SetSkip((page - 1) * 20)
	cursor, err := col.Find(ctx, condition, opti)
	if err != nil {
		return results, false
	}
	for cursor.Next(ctx) {
		var s models.User
		cursor.Decode(&s)
		s.Password = ""
		results = append(results, &s)
	}
	return results, true
}
