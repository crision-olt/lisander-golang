package routes

import (
	"net/http"
	"time"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertRelation performs the registration of the relationship between users*/
func InsertRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The parameter ID is required", http.StatusBadRequest)
		return
	}

	var t models.RelationInsert
	t.UserID = UserID
	t.UserRelationID = ID
	t.Date = time.Now()

	status, err := database.InsertRelation(t)
	if err != nil {
		http.Error(w, "An error occurred when trying to insert relationship "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "It has not been possible to insert the relation"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
