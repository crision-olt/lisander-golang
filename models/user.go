package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*User it is the model of user of the DataBase*/
type User struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name,omitempty"`
	Surnames        string             `bson:"surnames" json:"surnames,omitempty"`
	BirthDate       time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email           string             `bson:"email" json:"email"`
	UpdateDate      int                `bson:"updateDate" json:"updateDate"`
	UpdateDateToday int                `bson:"updateDateToday" json:"updateDateToday"`
	DateFrom        time.Time          `bson:"dateFrom" json:"dateFrom"`
	Password        string             `bson:"password" json:"password,omitempty"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempty"`
	Banner          string             `bson:"banner" json:"banner,omitempty"`
	Biography       string             `bson:"biography" json:"biography,omitempty"`
	Location        string             `bson:"location" json:"location,omitempty"`
	WebSite         string             `bson:"webSite" json:"webSite,omitempty"`
	Blocked         bool               `bson:"blocked" json:"blocked,omitempty"`
	Admin           int                `bson:"admin" json:"admin,omitempty"`
}
