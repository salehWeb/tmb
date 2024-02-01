package routes

import (
	"net/http"

	"github.com/salehWeb/chat-app/server/src/controllers/auth"
)


func HandelRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/sing-up/", auth.SingUp)
	mux.HandleFunc("/auth/login/", auth.Login)
	mux.HandleFunc("/auth/logout/", auth.Logout)

	return mux;
}

