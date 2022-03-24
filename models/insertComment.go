package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertComment structure of the comments*/
type InsertComment struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	TootID  string             `bson:"tootId" json:"tootId,omitempty"`
	UserID  string             `bson:"userId" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
	Hide    bool               `bson:"hide" json:"hide,omitempty"`
	Delete  bool               `bson:"delete" json:"delete,omitempty"`
}
