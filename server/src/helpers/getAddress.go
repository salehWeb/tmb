package helpers

import (
	"fmt"
	"os"
)

func GetAddress() string {
	var port string

	if len(os.Getenv("PORT")) > 1 {
		port = os.Getenv("PORT")
	}  else {
		port = "8080"
	}

	address := "127.0.0.1:" + port

	if os.Getenv("PRODUCTION") == "true" {
		fmt.Printf("App Running In Production, Port: %s\n\n", port)
		address = ":" + port
	} else {
		fmt.Printf("App listen in \"http://localhost:%s\"\n\n", port)
	}

	return address;
}
