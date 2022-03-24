package models

import (
	"time"
)

/*Report structure of reports*/
type Report struct {
	UserID      string    `bson:"userId" json:"userId,omitempty"`
	Description string    `bson:"description" json:"description,omitempty"`
	Date        time.Time `bson:"date" json:"date,omitempty"`
}
