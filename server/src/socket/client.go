package socket

import (
	"fmt"
	"log"

	"os"

	"github.com/gorilla/websocket"
	"github.com/salehWeb/chat-app/server/src/helpers"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type int    `json:"type"`
	Body string `json:"body"`
}

type Data struct {
	Type string `json:"type"`
	Body []byte `json:"body"`
}

func (c *Client) Read() {

	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, bytes, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			break
		}

		fmt.Printf("Message Type: %d", messageType)

		if messageType == 1 {
			message := Message{Type: messageType, Body: string(bytes)}
			c.Pool.Broadcast <- message
			fmt.Printf("Message Received: %+v\n", message)
			continue
		}

		if messageType == 2 {

			fmt.Println("Its Audio Pro")
			message := Message{Type: messageType, Body: "Its Audio Pro"}
			c.Pool.Broadcast <- message

			f, err := os.Create(fmt.Sprintf("%s.weba", helpers.UUID()))

			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(bytes)
			f.Write(bytes)
			fmt.Println(f.Name())
			f.Close()

			continue
		}

	}
}
