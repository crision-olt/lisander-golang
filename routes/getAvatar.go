package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/crision98/lisander-golang-backend/database"
)

/*GetAvatar allows to get Avatar from a profile to HTTP*/
func GetAvatar(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Must send the parameter ID", http.StatusBadRequest)
		return
	}

	profile, err := database.GetUser(ID)
	if err != nil {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Image not found", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Failed to copy the image", http.StatusBadRequest)
	}

}
