package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/crision98/lisander-golang-backend/models"
)

/*GetDate modify database update date*/
func GetDate(ID string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var dates models.User

	db := MongoCN.Database("lisander")
	col := db.Collection("user")
	objID, _ := primitive.ObjectIDFromHex(ID)

	err := col.FindOne(ctx, bson.M{"_id": bson.M{"$eq": objID}}).Decode(&dates)
	if err != nil {
		return dates, err
	}
	if dates.UpdateDate == 0 {
		dates.DateFrom = time.Now()
	}

	return dates, nil
}
