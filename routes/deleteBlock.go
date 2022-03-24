package routes

import (
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*DeleteBlock performs the deletion of block on a user*/
func DeleteBlock(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Block
	t.UserID = UserID
	t.BlockedUserID = ID

	status, err := database.DeleteBlock(t)
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
