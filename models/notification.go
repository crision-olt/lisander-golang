package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

/*Notification it is the model of user of the DataBase*/
type Notification struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TypeID        int                `bson:"typeId,omitempty" json:"typeId"`
	UserID        string             `bson:"userId,omitempty" json:"userId,omitempty"`
	FromUserID    string             `bson:"fromUserId,omitempty" json:"fromUserId,omitempty"`
	FromUserName  string             `bson:"fromUserName,omitempty" json:"fromUserName,omitempty"`
	FromTootID    string             `bson:"fromTootId,omitempty" json:"fromTootId,omitempty"`
	FromCommentID string             `bson:"fromCommentId,omitempty" json:"fromCommentId,omitempty"`
	IdDelete      string             `bson:"idDelete,omitempty" json:"idDelete,omitempty"`
	Watched       string             `bson:"watched,omitempty" json:"watched,omitempty"`
	Date          time.Time          `bson:"date" json:"date,omitempty"`
}
