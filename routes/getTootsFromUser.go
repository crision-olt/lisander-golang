package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crision98/lisander-golang-backend/database"
)

/*GetTootsFromUser read the toots*/
func GetTootsFromUser(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "You must send the parameter id", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "You must send the parameter page", http.StatusBadRequest)
		return
	}
	page1, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "You must send the parameter page with a value greater than 0", http.StatusBadRequest)
		return
	}
	page2 := int64(page1)
	answer, correct := database.GetTootsFromUser(UserID, ID, page2)
	if !correct {
		http.Error(w, "Error while reading tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
