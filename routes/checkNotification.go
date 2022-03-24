package routes

import (
	"github.com/crision98/lisander-golang-backend/database"
	"net/http"
)

/*CheckNotification e*/
func CheckNotification(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter ID", http.StatusBadRequest)
		return
	}
	_, err := database.CheckNotification(ID)
	if err != nil {
		http.Error(w, "An error occurred when trying to search for the record "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
