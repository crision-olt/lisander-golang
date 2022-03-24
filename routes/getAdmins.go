package routes

import (
	"encoding/json"
	"github.com/crision98/lisander-golang-backend/database"
	"net/http"
)

/*GetAdmins shows a profile*/
func GetAdmins(w http.ResponseWriter, r *http.Request) {
	if Admin != 2 {
		http.Error(w, "You dont have permission to do that.", 400)
		return
	}
	result, status := database.GetAdmins()
	if !status {
		http.Error(w, "An error occurred when trying to search for the record ", 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
