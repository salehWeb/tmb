package config

import (
	"net/http"

	"github.com/salehWeb/chat-app/server/src/controllers"
)


func HandelRoutes() {
	http.HandleFunc("/api/sing-up/", controllers.SingUpControllers)
}

