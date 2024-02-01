package helpers

import (
	"crypto/rand"
	"fmt"
	"log"
)

func UUID() string {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)

	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x",
		bytes[0:4], bytes[4:6], bytes[6:8], bytes[8:10], bytes[10:])
}
