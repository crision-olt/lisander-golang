package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*CountCommentFromToot returns count of comments of a toot*/
func CountCommentsFromToot(w http.ResponseWriter, r *http.Request) {
	TootID := r.URL.Query().Get("tootId")
	if len(TootID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}
	count, err := database.CountCommentsFromToot(UserID, TootID)
	if err != nil {
		http.Error(w, "An error occurred when trying to search for the record "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(count)
}
