package models

/*ChangePassword it is the model of user of the DataBase*/
type ChangePassword struct {
	Password       string `json:"password"`
	RepeatPassword string `json:"repeatPassword"`
	NewPassword    string `json:"newPassword"`
}
