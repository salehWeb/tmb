package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/salehWeb/chat-app/server/src/helpers"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upGrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upGrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}


func serveWS(pool *Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("websocket endpoint reached")

	conn, err := Upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &Client{
		ID: helpers.UUID(),
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client

	client.Read()
}

func UseSocket() *http.ServeMux  {
    pool := NewPool()
    go pool.Start()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { serveWS(pool, w, r) })

	return mux
}
