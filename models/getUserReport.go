package models

/*GetUserReport structure to follow to make petition to report some user.*/
type GetUserReport struct {
	ReportedUserID string `bson:"reportedUser" json:"reportedUser,omitempty"`
	Description    string `bson:"description" json:"description,omitempty"`
}
