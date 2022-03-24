package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crision98/lisander-golang-backend/database"
)

/*GetReports get the reports maded on users*/
func GetReports(w http.ResponseWriter, r *http.Request) {
	if Admin == 0 {
		http.Error(w, "You don't have permission", http.StatusBadRequest)
		return
	}
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Must send the parameter page", http.StatusBadRequest)
		return
	}
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Must send the parameter page like a int greater than 0", http.StatusBadRequest)
		return
	}
	answer, correct := database.GetReports(int64(page))
	if !correct {
		http.Error(w, "Error while reading the toots", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(answer)
}
