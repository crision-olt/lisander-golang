package models

import (
	"time"
)

/*RelationInsert is the structure for save the relation of the user with another user */
type RelationInsert struct {
	UserID         string    `bson:"userId" json:"userId"`
	UserRelationID string    `bson:"userRelationId" json:"userRelationId"`
	Date           time.Time `bson:"date" json:"date"`
}
