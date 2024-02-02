package socket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/salehWeb/chat-app/server/src/helpers"
)


func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	// handshake 101
	upGrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}

	// set headers and return the socket connection
	conn, err := upGrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}


func serveWS(pool *Pool, w http.ResponseWriter, r *http.Request) {
	// runs for every client at first handshake
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

func UseSocket() {
    pool := NewPool()
    go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { serveWS(pool, w, r) })
}
