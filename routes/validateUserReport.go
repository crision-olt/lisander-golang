package routes

import (
	"encoding/json"
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*ValidateUserReport ban the user that is provided in id is the banned slot is in true*/
func ValidateUserReport(w http.ResponseWriter, r *http.Request) {
	if Admin == 0 {
		http.Error(w, "You don't have permission", http.StatusBadRequest)
		return
	}

	ID := r.URL.Query().Get("id")
	banned := r.URL.Query().Get("banned")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}
	profile, err := database.ValidateUserReport(ID, banned, Admin > 1)
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
