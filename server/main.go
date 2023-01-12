package main

import (
	"github.com/salehWeb/chat-app/server/src/config"
	"github.com/salehWeb/chat-app/server/src/initializers"
	"github.com/salehWeb/chat-app/server/src/routes"
)

func main() {
	initializers.GetENV()
	go initializers.ConnectToDatabase()
	go config.UseClient()
	go routes.HandelRoutes()
	initializers.Listen()
}
