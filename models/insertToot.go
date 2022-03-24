package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertToot structure of our toots when we record them*/
type InsertToot struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"userId" json:"userId,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
	Hide    bool               `bson:"hide" json:"hide,omitempty"`
	Deleted bool               `bson:"deleted" json:"deleted,omitempty"`
}
