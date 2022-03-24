package database

import (
	"context"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertBlock inserts relation in DB*/
func InsertBlock(t models.Block) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("block")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}
	DeleteRelation(models.Relation{UserID: t.BlockedUserID, UserRelationID: t.UserID})
	DeleteRelation(models.Relation{UserID: t.UserID, UserRelationID: t.BlockedUserID})
	return true, nil
}
