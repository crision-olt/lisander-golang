package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*GetDate used to get the date from where gets the notifications and publications.*/
func GetDate(w http.ResponseWriter, r *http.Request) {
	var answer models.User
	answer, err := database.GetDate(UserID)
	if err != nil {
		http.Error(w, "An error occurred when trying to search for the record "+err.Error(), 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer.DateFrom)
}
