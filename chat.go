package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client represents a connected user
type Client struct {
	conn *websocket.Conn
}

// Message represents a message from a user
type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Text      string `json:"text"`
}

var clients = make(map[string]*Client)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	client := &Client{conn: conn}
	username := r.URL.Query().Get("username")

	if username == "" {
		log.Println("Username not provided")
		return
	}

	clients[username] = client
	log.Printf("User %s connected\n", username)

	for {
		var message Message
		err := conn.ReadJSON(&message)
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			break
		}

		log.Printf("Received message from %s to %s: %s\n", message.Sender, message.Recipient, message.Text)

		recipient, ok := clients[message.Recipient]
		if !ok {
			log.Printf("User %s not found\n", message.Recipient)
			continue
		}

		recipient.conn.WriteJSON(message)
	}
}

func main() {
	http.HandleFunc("/", handleConnections)

	port := 8080
	fmt.Printf("Server is running on :%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
