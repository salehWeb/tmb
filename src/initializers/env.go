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
		item := strings.TrimSpace(string(keysValues[i]))

		if len(item) < 4 {
			return
		}

		keyValue := strings.Split(item, "=\"")
		os.Setenv(keyValue[0], strings.Split(keyValue[1], "\"")[0])
	}

}
