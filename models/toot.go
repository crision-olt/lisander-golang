package models

/*Toot captures the body, the message that comes to us*/
type Toot struct {
	Message string `bson:"message" json:"message"`
}
