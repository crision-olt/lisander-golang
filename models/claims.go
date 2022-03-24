package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim is the structure used to process the JWT */
type Claim struct {
	Email   string             `json:"email"`
	Admin   int                `json:"admin"`
	Avatar  string             `json:"avatar,omitempty"`
	Blocked bool               `json:"blocked,omitempty"`
	ID      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
