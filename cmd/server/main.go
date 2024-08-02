package main

import (
	"bufio"
	"fmt"
	"go-chat-app/config"
	"go-chat-app/models"
	"log"
	"net"
	"strings"
	"sync"
)

var clients = make(map[net.Conn]string)
var messages = make(chan models.Message)
var mutex = &sync.Mutex{}

func main() {
	config.LoadConfig()
	server, err := net.Listen("tcp", config.GetConfig().ServerAddress+":"+config.GetConfig().ServerPort)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer server.Close()

	fmt.Printf("Server started on %s:%s\n", config.GetConfig().ServerAddress, config.GetConfig().ServerPort)

	go handleMessages()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}
		go handleConnection(conn)
	}

}

func handleMessages() {
	for msg := range messages {
		broadcastMessage(msg)
	}
}

func broadcastMessage(msg models.Message) {
	mutex.Lock()
	defer mutex.Unlock()

	for client := range clients {
		_, err := fmt.Fprintf(client, "%s: %s\n", msg.Username, msg.Text)
		if err != nil {
			log.Printf("Error sending message to client: %v", err)
			client.Close()
			delete(clients, client)
		}
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Fprint(conn, "Enter your username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Failed to read username: %v", err)
		return
	}
	username = strings.TrimSpace(username)

	mutex.Lock()
	clients[conn] = username
	mutex.Unlock()

	fmt.Fprintf(conn, "Welcome, %s!\n", username)

	for {
		msgText, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Failed to read message: %v", err)
			break
		}
		msg := models.Message{
			Username: username,
			Text:     strings.TrimSpace(msgText),
		}
		messages <- msg
	}

	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()

	fmt.Printf("%s has left the chat\n", username)
}
