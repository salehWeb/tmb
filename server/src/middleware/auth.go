package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/salehWeb/chat-app/server/src/helpers"
)

func Authorized(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("User is UnAuthorized"))
			return
		}

		token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("User is UnAuthorized"))
			}

			return helpers.GetSecretKey(), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("User is UnAuthorized"))
			return
		}

		if token.Valid {
			next(w, r)
		}
	})
}
