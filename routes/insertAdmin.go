package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertAdmin is the function to create the user registry in the DB*/
func InsertAdmin(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if Admin != 2 {
		http.Error(w, "You dont have permission to do that "+err.Error(), 400)
		return
	}
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "User email is required ", 400)
		return
	}
	t.Password = "admin01"
	t.UpdateDate = 0
	t.UpdateDateToday = t.UpdateDate
	_, found, _ := database.UserAlreadyExist(t.Email)
	if found {
		http.Error(w, "There is already a registered user with this email ", 400)
		return
	}
	t.DateFrom = time.Now()
	t.Blocked = false
	t.Admin = 1

	_, status, err := database.InsertAdmin(t)
	if err != nil {
		http.Error(w, "An error occurred while trying to register "+err.Error(), 400)
	}
	if !status {
		http.Error(w, "User registration failed to be inserted", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
