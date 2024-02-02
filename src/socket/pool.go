package socket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {

		case newClient := <-pool.Register:
			pool.Clients[newClient] = true
			text := fmt.Sprintf("New User Joined total %d", len(pool.Clients))
			message := Message{Type: 1, Body: text, ClientId: newClient.ID}

			for client := range pool.Clients {
				if client.ID == newClient.ID {
					continue
				}
				client.Conn.WriteJSON(message)
			}

		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			message := Message{Type: 1, Body: "User Disconnected...", ClientId: client.ID}

			for client := range pool.Clients {
				client.Conn.WriteJSON(message)
			}

		case message := <-pool.Broadcast:
			for client := range pool.Clients {
				if message.ClientId == client.ID {
					continue
				}

				err := client.Conn.WriteJSON(message)
				if err != nil {
					pool.Unregister <- client
					client.Conn.Close()
				}
			}

		}

	}
}
