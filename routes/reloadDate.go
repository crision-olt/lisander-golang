package routes

import (
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*ReloadDate used to update the date from where gets the notifications and publications.*/
func ReloadDate(w http.ResponseWriter, r *http.Request) {
	status, err := database.ReloadDate(UserID)
	if err != nil {
		http.Error(w, "An error occurred when trying to modify the record. Retry again "+err.Error(), http.StatusBadRequest)
		return
	}
	if !status {
		http.Error(w, "User registration could not be modified "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
