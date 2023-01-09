package helpers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey []byte

func NewToken() string {



	if len(os.Getenv("SECRET_KEY")) > 1 {
		secretKey = []byte(os.Getenv("SECRET_KEY"))
	} else {
		fmt.Println("No SECRET KEY Found For Jwt Will Use \"Hello World\" as SECRET KEY")
		secretKey = []byte("Hello World")
	}

	tokenConfig := jwt.New(jwt.SigningMethodHS256)

	claims := tokenConfig.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(time.Now().Year())).Unix()

	token, err := tokenConfig.SignedString(secretKey)

	if err != nil {
		panic(err)
	}

	return token
}


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

			return secretKey, nil
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
