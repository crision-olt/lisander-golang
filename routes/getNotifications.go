package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crision98/lisander-golang-backend/database"
)

/*GetNotifications read the toots*/
func GetNotifications(w http.ResponseWriter, r *http.Request) {
	page1, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the parameter page with a value greater than 0", http.StatusBadRequest)
		return
	}
	page2 := int64(page1)
	answer, correct := database.GetNotifications(UserID, page2)
	if !correct {
		http.Error(w, "Error while reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
