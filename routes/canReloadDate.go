package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*CanReloadDate used to check if can update the date from where gets the notifications and publications.*/
func CanReloadDate(w http.ResponseWriter, r *http.Request) {
	var answer models.AnswerCheck
	status, err := database.CanReloadDate(UserID)
	if err != nil || !status {
		answer.Status = false
	} else {
		answer.Status = true
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
