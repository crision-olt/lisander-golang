package routes

import (
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*DeleteRelation performs the deletion of the relationship between users*/
func DeleteRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = UserID
	t.UserRelationID = ID

	status, err := database.DeleteRelation(t)
	if err != nil {
		http.Error(w, "An error occurred when trying to delete the relationship"+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "I can't erase the relationship", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
