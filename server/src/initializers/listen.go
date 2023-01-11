package initializers

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func Listen() {
	var port string

	if len(os.Getenv("PORT")) > 1 {
		port = os.Getenv("PORT")
	}  else {
		port = "8080"
	}

	if os.Getenv("PRODUCTION") == "true" {
		fmt.Printf("App Running In Production, Port: %s\n\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	} else {
		fmt.Printf("App listen in \"http://localhost:%s\"\n\n", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%s", port), nil))
	}
}
