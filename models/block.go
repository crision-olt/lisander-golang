package models

/*Block structure of the blocks*/
type Block struct {
	UserID        string `bson:"userId" json:"userId,omitempty"`
	BlockedUserID string `bson:"blockedUserId" json:"blockedUserId,omitempty"`
}
