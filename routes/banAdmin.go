package routes

import (
	"encoding/json"
	"github.com/crision98/lisander-golang-backend/database"
	"net/http"
)

/*banAdmin ban the user that is provided in id is the banned slot is in true*/
func BanAdmin(w http.ResponseWriter, r *http.Request) {
	if Admin <= 1 {
		http.Error(w, "You don't have permission", http.StatusBadRequest)
		return
	}

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}
	profile, err := database.BanAdmin(ID)
	if !profile {
		http.Error(w, "You don't have permission", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "An error occurred when trying to search for the record "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
