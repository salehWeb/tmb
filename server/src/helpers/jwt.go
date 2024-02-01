package helpers

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)


func NewToken(id uint) string {
	tokenConfig := jwt.New(jwt.SigningMethodHS256)

	claims := tokenConfig.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Duration(time.Now().Year())).Unix()
	claims["id"] = id

	token, err := tokenConfig.SignedString(GetSecretKey())

	if err != nil {
		panic(err)
	}

	return token
}



func GetSecretKey() []byte {
	var secretKey []byte

	if len(os.Getenv("SECRET_KEY")) > 1 {
		secretKey = []byte(os.Getenv("SECRET_KEY"))
	} else {
		fmt.Println("No SECRET KEY Found For Jwt Will Use \"Hello World\" as SECRET KEY")
		secretKey = []byte("Hello World")
	}

	return secretKey
}
