package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*CheckRelation check if exist a relation between 2 users*/
func CheckRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	var answer models.AnswerCheck

	status, err := database.CheckRelation(t)
	if err != nil || !status {
		answer.Status = false
	} else {
		answer.Status = true
	}
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)

}
