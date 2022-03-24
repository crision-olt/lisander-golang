package models

/*Relations is the structure for save the relation of the user with another user */
type Relations struct {
	UserRelationID string `bson:"userRelationId" json:"userRelationId"`
}
