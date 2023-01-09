package routes

import (
	"net/http"

	"github.com/salehWeb/chat-app/server/controllers"
)

func HandelRoutes() {

	http.HandleFunc("/api/", controllers.HomeControllers)
	http.HandleFunc("/api/auth/", controllers.AuthControllers)

}
