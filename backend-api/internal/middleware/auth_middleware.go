package middleware

import (
	"fmt"
	"net/http"

	"github.com/expose443/real-time-forum-golang/backend-api/internal/jwt"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("i work")
		c, err := r.Cookie("jwt_token")
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		status, _, err := jwt.VerifyJWT(c.Value)
		fmt.Println(status, err)
		if status {
			fmt.Println(status)
			http.Redirect(w, r, "/home", http.StatusFound)
			return
		}
		fmt.Println("i still work")
		next.ServeHTTP(w, r)
	})
}
