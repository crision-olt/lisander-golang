package database

import (
	"github.com/crision98/lisander-golang-backend/models"
	"golang.org/x/crypto/bcrypt"
)

/*AttemptLogin performs the database login check*/
func AttemptLogin(email string, password string) (models.User, bool) {
	usu, found, _ := UserAlreadyExist(email)
	if !found {
		return usu, false
	}
	passwordBytes := []byte(password)
	passworddatabase := []byte(usu.Password)
	err := bcrypt.CompareHashAndPassword(passworddatabase, passwordBytes)
	if err != nil {
		return usu, false
	}
	return usu, true
}
