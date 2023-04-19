package middle

import (
	"fmt"
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/jwt"
)

func AuthRedirectMiddleware(next http.HandlerFunc) http.HandlerFunc {
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

func RequireAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt_token")
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		status, _, err := jwt.VerifyJWT(c.Value)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			fmt.Println(err)
			return

		}
		if !status {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
