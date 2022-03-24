package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*FindUsersByEmail reads the list of users by email*/
func FindUsersByEmail(w http.ResponseWriter, r *http.Request) {
	results, status := database.FindUserByEmail(UserID)
	if !status {
		http.Error(w, "Error while reading the users", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(results)

}
