package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*GetBlockedUsers allows search users by his name*/
func GetBlockedUsers(userID string, page int64) ([]*models.User, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	opti := options.Find()
	var results []*models.User
	userIds, _ := GetMyBlocks(userID)

	db := MongoCN.Database("lisander")
	col := db.Collection("user")

	condition := bson.M{
		"_id": bson.M{
			"$in": userIds,
		},
	}

	opti.SetLimit(20)
	opti.SetSkip((page - 1) * 20)
	cursor, err := col.Find(ctx, condition, opti)
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
