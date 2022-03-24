package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/crision98/lisander-golang-backend/database"
)

/*GetBanner allows to get Banner from a profile to HTTP*/
func GetBanner(w http.ResponseWriter, r *http.Request) {

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

	OpenFile, err := os.Open("uploads/banners/" + profile.Banner)
	if err != nil {
		http.Error(w, "Banner not found", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error while copying the banner", http.StatusBadRequest)
	}

}
