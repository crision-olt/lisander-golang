package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertUserReport insert a report of a user.*/
func InsertUserReport(w http.ResponseWriter, r *http.Request) {
	var t models.GetUserReport
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}

	status, err := database.InsertUserReport(t, UserID)
	if err != nil {
		http.Error(w, "An error occurred while trying to register "+err.Error(), 400)
	}
	if !status {
		http.Error(w, "User registration failed to be inserted", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
