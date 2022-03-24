package routes

import (
	"errors"
	"strings"

	"github.com/crision98/lisander-golang-backend/database"

	"github.com/crision98/lisander-golang-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*Email value of Email used in all EndPoints*/
var Email string

/*UserID is the ID returned from the model, used in all EndPoints*/
var UserID string

/*Admin e*/
var Admin int

/*ProcessToken process the token*/
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("crazyforsnowboards")
	claims := &models.Claim{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid format of token")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, found, _ := database.UserAlreadyExist(claims.Email)
		if found {
			Admin = claims.Admin
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid Token")
	}

	return claims, false, string(""), err
}
