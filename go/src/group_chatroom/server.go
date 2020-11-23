package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/websocket"
)

// ClientManager struct
type ClientManager struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	connect    chan *Client
	disconnect chan *Client
}

// Client struct
type Client struct {
	id       string
	username string
	socket   *websocket.Conn
	send     chan []byte
}

// Message will convert to json
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

var manager = ClientManager{
	clients:    make(map[*Client]bool),
	broadcast:  make(chan []byte),
	connect:    make(chan *Client),
	disconnect: make(chan *Client),
}

func (manager *ClientManager) start() {
	for {
		select {
		case conn := <-manager.connect:
			manager.clients[conn] = true
			content := fmt.Sprintf("Welcone %v join the chatroom", conn.username)
			jsonMessage, _ := json.Marshal(&Message{Content: content, Sender: "system"})
			manager.send(jsonMessage)
		case conn := <-manager.disconnect:
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				content := fmt.Sprintf("%v has lefted the chatroom", conn.username)
				jsonMessage, _ := json.Marshal(&Message{
					Content: content,
					Sender:  "system"})
				manager.send(jsonMessage)
			}
		case message := <-manager.broadcast:
			for conn := range manager.clients {
				select {
				case conn.send <- message:
				default:
					close(conn.send)
					delete(manager.clients, conn)
				}
			}
		}
	}
}

func (manager *ClientManager) send(message []byte) {
	for conn := range manager.clients {
		conn.send <- message
		fmt.Println("send conn is ", conn)
	}
}

func (c *Client) read() {
	defer func() {
		manager.disconnect <- c
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			manager.disconnect <- c
			c.socket.Close()
			break
		}
		jsonMessage, _ := json.Marshal(&Message{Sender: c.username, Content: string(message)})
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

func wsHandler(res http.ResponseWriter, req *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	var username string
	params := req.URL.Query()
	if params["username"] != nil {
		username = params["username"][0]
	} else {
		username = "visitor"
	}
	client := &Client{
		id: uuid.Must(uuid.NewV4()).String(), socket: conn, send: make(chan []byte),
		username: username}

	manager.connect <- client

	go client.read()
	go client.write()
}

func main() {
	fmt.Println("Start a group chatroom")
	go manager.start()
	http.HandleFunc("/chat", wsHandler)
	http.ListenAndServe(":9000", nil)
}
