package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*ChangePassword modify the user's profile*/
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var t models.ChangePassword
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrect Data"+err.Error(), http.StatusBadRequest)
		return
	}
	_, exist := database.AttemptLogin(Email, t.Password)
	if !exist {
		http.Error(w, "Invalid Password ", 400)
		return
	}
	status, err := database.ChangePassword(t, UserID)
	if err != nil {
		http.Error(w, "An error occurred when trying to modify the record. Retry again "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "User registration could not be modified "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
