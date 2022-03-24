package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*CountTootsFromUser return the count of the toot maked by the user*/
func CountTootsFromUser(w http.ResponseWriter, r *http.Request) {
	count, err := database.CountTootsFromUser(UserID)
	if err != nil {
		http.Error(w, "An error occurred when trying to search for the record "+err.Error(), 400)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(count)
}
