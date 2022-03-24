package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
)

/*UploadBanner allows to upload our Banner*/
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archive string = "uploads/banners/" + UserID + "." + extension

	f, err := os.OpenFile(archive, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error while uploading the banner ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error while copying the banner ! "+err.Error(), http.StatusBadRequest)
		return
	}
	var user models.User

	user.Banner = UserID + "." + extension
	status, err := database.ModifyUser(user, UserID)
	if err != nil || !status {
		http.Error(w, "Error while recording the banner ! "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
