package routes

import (
	"encoding/json"
	"github.com/crision98/lisander-golang-backend/database"
	"net/http"
)

func CountNotifications(w http.ResponseWriter, r *http.Request) {

	count, err := database.CountNotifications(UserID)
	if err != nil {
		http.Error(w, "Error while reading the users", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(count)
}
