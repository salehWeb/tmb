package config

import (
	"net/http"

	"github.com/salehWeb/chat-app/server/src/controllers"
)


func HandelRoutes() {
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)
	http.HandleFunc("/api/sing-up/", controllers.SingUpControllers)

}
