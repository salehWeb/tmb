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

	var addr string = "127.0.0.1:" + port




	if os.Getenv("PRODUCTION") == "true" {
		fmt.Printf("App Running In Production, Port: %s\n\n", port)
		addr = ":" + port
	} else {
		fmt.Printf("App listen in \"http://localhost:%s\"\n\n", port)
	}



	log.Fatal(http.ListenAndServe(addr, nil))
}
