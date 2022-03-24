package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*UserReport structure of the user reports*/
type UserReport struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	IDReportedUser string             `bson:"idReportedUser" json:"idReportedUser"`
	ReportedUser   User               `bson:"reportedUser" json:"reportedUser,omitempty"`
	Reports        []Report           `bson:"reports" json:"reports,omitempty"`
	Solved         bool               `bson:"solved" json:"solved,omitempty"`
	Date           time.Time          `bson:"date" json:"date,omitempty"`
}
