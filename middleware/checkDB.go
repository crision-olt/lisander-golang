package middleware

import (
	"net/http"

	"github.com/crision98/lisander-golang-backend/database"
)

/*CheckDB returns an anonymous function, which will check that the server is connected and will tell the HTTP server to return what we have asked for*/
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(w, "Connection lost with the DataBase", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
