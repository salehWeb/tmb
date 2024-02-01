package main

import (
	"log"
	"net/http"

	"github.com/salehWeb/chat-app/server/src/helpers"
	"github.com/salehWeb/chat-app/server/src/initializers"
	"github.com/salehWeb/chat-app/server/src/middleware"
	"github.com/salehWeb/chat-app/server/src/routes"
	"github.com/salehWeb/chat-app/server/src/socket"
)

func main() {
	initializers.GetENV()
	initializers.InitClient()

	http.Handle("/", middleware.GzipHandler(routes.HandelClient()))
	http.Handle("/assets/", middleware.GzipHandler(http.FileServer(http.Dir("./dist"))))
	http.Handle("/api/", middleware.GzipHandler(routes.HandelRoutes()))

	socket.UseSocket()

	log.Fatal(http.ListenAndServe(helpers.GetAddress(), nil))
}
