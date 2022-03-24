package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertRelation inserts relation in DB*/
func InsertRelation(t models.RelationInsert) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	CreateNotification(models.Notification{TypeID: 1, UserID: t.UserRelationID, FromUserID: t.UserID})

	return true, nil
}
