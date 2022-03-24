package database

import (
	"context"
	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*CreateNotification records toot in DB*/
func CreateNotification(t models.Notification) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	u, _ := GetUser(t.FromUserID)
	t.FromUserName = u.Name + " " + u.Surnames
	t.Watched = "true"
	t.Date = time.Now()
	db := MongoCN.Database("lisander")
	col := db.Collection("notification")
	result, err := col.InsertOne(ctx, t)
	if err != nil {
		return string(""), false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
