package initializers

import (
	"fmt"
	"os"
	"strings"
)

func GetENV() {

	v, err := os.ReadFile("./.env")

	if err != nil {
		fmt.Printf("No .env file In The Root Directory\n\n")
		return
	}

	var str string = strings.TrimSpace(string(v))

	keysValues := strings.Split(str, "\n")

	for i := 0; i < len(keysValues); i++ {
		s := strings.Split(keysValues[i], "=\"")
		os.Setenv(s[0], strings.Split(s[1], "\"")[0])
	}

}
