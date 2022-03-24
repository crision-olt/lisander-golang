package routes

import (
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertBlock  block a user*/
func InsertBlock(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "The parameter ID is required", http.StatusBadRequest)
		return
	}

	var t models.Block
	t.UserID = UserID
	t.BlockedUserID = ID

	status, err := database.InsertBlock(t)
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
