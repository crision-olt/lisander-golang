package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertToot  to create a toot*/
func InsertToot(w http.ResponseWriter, r *http.Request) {
	var message models.Toot
	json.NewDecoder(r.Body).Decode(&message)

	record := models.InsertToot{
		UserID:  UserID,
		Message: message.Message,
		Date:    time.Now(),
		Hide:    false,
		Deleted: false,
	}

	_, status, err := database.InsertToot(record)
	if err != nil {
		http.Error(w, "An error occurred when trying to insert the record, retry again "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "The toot has not been inserted", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
