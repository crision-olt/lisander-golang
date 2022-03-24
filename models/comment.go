package models

/*Comment structure of comment*/
type Comment struct {
	TootID  string `bson:"tootId" json:"tootId"`
	Message string `bson:"message" json:"message"`
}
