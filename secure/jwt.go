package secure

import (
	"time"

	"github.com/crision98/lisander-golang-backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GenerateJWT generates the encrypted token with JWT*/
func GenerateJWT(t models.User) (string, error) {
	miClave := []byte("crazyforsnowboards")
	payload := jwt.MapClaims{
		"email": t.Email,
		"admin": t.Admin,
		"_id":   t.ID.Hex(),
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
