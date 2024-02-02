package routes

import (
	"net/http"

	"github.com/salehWeb/chat-app/server/src/controllers/auth"
)

var clientPaths = map[string]bool{
	"/login":   true,
	"/sign-up": true,
	"/chat":    true,
	"/":        true,
}

func HandelRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth/sing-up/", auth.SingUp)
	mux.HandleFunc("/api/auth/login/", auth.Login)
	mux.HandleFunc("/api/auth/logout/", auth.Logout)
	mux.HandleFunc("/api/auth/", auth.Test)
	return mux;
}

func HandelClient() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if clientPaths[r.URL.Path] {
			http.ServeFile(w, r, "./dist/index.html")
		}
	})
}
