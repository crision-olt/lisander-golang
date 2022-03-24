package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/crision98/lisander-golang-backend/database"
	"github.com/crision98/lisander-golang-backend/models"
	"github.com/crision98/lisander-golang-backend/secure"
)

/*Login log in*/
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid User and/or Password "+err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		http.Error(w, "The user's email is required ", 400)
		return
	}
	document, exist := database.AttemptLogin(t.Email, t.Password)
	if !exist {
		http.Error(w, "Invalid User and/or Password ", 400)
		return
	}
	if document.Blocked {
		http.Error(w, "This account is banned", http.StatusMethodNotAllowed)
		return
	}
	jwtKey, err := secure.GenerateJWT(document)
	if err != nil {
		http.Error(w, "An error occurred when trying to generate the corresponding token "+err.Error(), 400)
		return
	}
	resp := models.AnswerLogin{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
