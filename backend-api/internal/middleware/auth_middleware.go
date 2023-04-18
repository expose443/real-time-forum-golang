package middleware

import (
	"fmt"
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt_token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		status, _, err := jwt.VerifyJWT(c.Value)
		if err != nil {
			fmt.Println(err)
		}
		if status {
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
