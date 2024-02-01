package chat

import (
	"fmt"
	"net/http"
)

func CreateChat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chat Created Successfully")
}
