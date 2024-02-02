package socket

import "github.com/gorilla/websocket"

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

type Message struct {
	Type     int    `json:"type"`
	Body     string `json:"body"`
	ClientId string `json:"clientId"`
}

func (c *Client) Read() {
	for {
		_, bytes, err := c.Conn.ReadMessage()

		if err != nil {
			c.Pool.Unregister <- c
			c.Conn.Close()
			break
		}

		message := Message{Type: 1, Body: string(bytes), ClientId: c.ID}
		c.Pool.Broadcast <- message
	}
}
