package main

import (
	"fmt"

	"github.com/salehWeb/chat-app/server/initializers"
)

func init() {
	fmt.Printf("Hello World 1\n\n")
	initializers.GetENV()
	db := initializers.ConnectToDatabase(string(initializers.ENV["FUCK"]))
	fmt.Println(db)
}

func main() {
	fmt.Printf("Hello World 2\n\n")
}
