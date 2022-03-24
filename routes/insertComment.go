package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*InsertComment  creates a comment on the toot provided*/
func InsertComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment
	json.NewDecoder(r.Body).Decode(&comment)

	insert := models.InsertComment{
		TootID:  comment.TootID,
		UserID:  UserID,
		Message: comment.Message,
		Date:    time.Now(),
		Hide:    false,
	}

	_, status, err := database.InsertComment(insert)
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
