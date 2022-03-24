package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crision98/lisander-golang-backend/database"
)

/*GetBlockedUsers get users blocked*/
func GetBlockedUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Must send the parameter page as integer greater than 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	result, status := database.GetBlockedUsers(UserID, pag)
	if !status {
		http.Error(w, "Error while reading the users", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}
