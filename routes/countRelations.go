package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*CountRelations return the count of followers and followings of a user*/
func CountRelations(w http.ResponseWriter, r *http.Request) {
	typeUser := r.URL.Query().Get("type")

	count, err := database.CountRelations(UserID, typeUser)
	if err != nil {
		http.Error(w, "Error while reading the users", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(count)
}
