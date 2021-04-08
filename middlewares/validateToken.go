package middlewares

import (
	"net/http"

	"github.com/DalinarLG/blog/routers"
)

func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, err := routers.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "M error proccesing token "+err.Error(), 400)
			return
		}
		next.ServeHTTP(w, r)
	}
}
