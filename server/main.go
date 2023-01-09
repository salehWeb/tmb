package main

import (
	"fmt"

	"github.com/salehWeb/chat-app/server/helpers"
	"github.com/salehWeb/chat-app/server/initializers"
	"github.com/salehWeb/chat-app/server/routes"
)

func main() {
	initializers.GetENV()
	initializers.ConnectToDatabase()
	fmt.Println(helpers.NewToken())

	routes.HandelRoutes()
	initializers.Listen()
}
