package middlewares

import (
	"net/http"

	"github.com/DalinarLG/blog/db"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.CheckDB() {
			http.Error(w, "Database error ", 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
