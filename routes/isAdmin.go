package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*IsAdmin used to check if can update the date from where gets the notifications and publications.*/
func IsAdmin(w http.ResponseWriter, r *http.Request) {
	status, err := database.IsAdmin(Admin)
	if err != nil {
		http.Error(w, "Error in received data "+err.Error(), 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.AnswerCheck{Status: status})
}
