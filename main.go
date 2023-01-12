package main

import (
	"github.com/salehWeb/chat-app/server/src/config"
	"github.com/salehWeb/chat-app/server/src/initializers"
)

func main() {
	initializers.GetENV()
	// initializers.ConnectToDatabase()
	config.UseClient()
	config.HandelRoutes()
	initializers.Listen()
}
