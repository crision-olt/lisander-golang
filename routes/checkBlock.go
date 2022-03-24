package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*CheckBlock check if exist a block between 2 users*/
func CheckBlock(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Block
	t.UserID = UserID
	t.BlockedUserID = ID

	var answer models.AnswerCheck

	status, err := database.CheckBlock(t)
	if err != nil || !status {
		answer.Status = false
	} else {
		answer.Status = true
	}
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)

}
