package config

import (
	"fmt"
	"net/http"
	"os"
)


var ClientRoutes [2]string = [2]string{"/", "/login/"}

func UseClient() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	for i := 0; i < len(ClientRoutes); i++ {
		http.HandleFunc(ClientRoutes[i], Serve)
	}
}


func Serve(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		bytes, err := os.ReadFile("./dist/index.html")

		if err != nil {
			fmt.Println("No File Found in ./dist/index.html")
			w.WriteHeader(http.StatusNotFound)
			w.Header().Set("Content-Type", "text/html")
			w.Write([]byte("Not Found"))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text/html")
		w.Write(bytes)
	}
}

