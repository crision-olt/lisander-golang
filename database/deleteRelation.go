package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*DeleteRelation deletes a relation between users*/
func DeleteRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("relation")
	filter := bson.M{
		"userId":         t.UserID,
		"userRelationId": t.UserRelationID,
	}
	var relation models.Relation
	col.FindOne(ctx, filter).Decode(relation)
	opt := bson.M{
		"typeId":     1,
		"userId":     t.UserRelationID,
		"fromUserId": t.UserID,
	}
	DeleteNotification(opt)
	_, err := col.DeleteOne(ctx, filter)
	if err != nil {
		return false, err
	}

	return true, nil
}
