package routes

import (
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*DeleteComment allows to delete a Toot*/
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}

	err := database.DeleteComment(ID, UserID)
	if err != nil {
		http.Error(w, "An error occurred when trying to delete the toot "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
