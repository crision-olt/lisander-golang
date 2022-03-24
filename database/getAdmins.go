package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*GetAdmins looking for profile in DB*/
func GetAdmins() ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("lisander")
	col := db.Collection("user")

	var results []*models.User
	condition := bson.M{
		"admin": 1,
	}
	opti := options.Find()
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
	if err != nil {
		return results, false
	}
	return results, true

}
