package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*Relation is the structure for save the relation of the user with another user */
type Relation struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID         string             `bson:"userId" json:"userId"`
	UserRelationID string             `bson:"userRelationId" json:"userRelationId"`
	Date           time.Time          `bson:"date" json:"date"`
}
