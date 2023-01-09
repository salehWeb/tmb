package controllers

import "net/http"


func HomeControllers(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte("<h1>Hello World From API@</h1>"))
		return
	}

	if r.Method == http.MethodPost {
		return
	}

	if r.Method == http.MethodPut {
		return
	}

	if r.Method == http.MethodDelete {
		return
	}

}
