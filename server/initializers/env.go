package initializers

import (
	"fmt"
	"os"
	"strings"
)

var ENV map[string]string = make(map[string]string)

func GetENV() {

	v, err := os.ReadFile("./.env")

	if err != nil {
		fmt.Println("No .env In The Root Directory")
	}

	var str string = strings.TrimSpace(string(v))

	keysValues := strings.Split(str, "\n")

	for i := 0; i < len(keysValues); i++ {
		s := strings.Split(keysValues[i], "=\"")
		ENV[s[0]] = strings.Split(s[1], "\"")[0]
	}

}
