package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"go-chat-app/config"
)

func main() {
	config.LoadConfig()

	conn, err := net.Dial("tcp", config.GetConfig().ServerAddress+":"+config.GetConfig().ServerPort)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	go readMessages(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text)
	}
}

func readMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read message: %v", err)
		}
		fmt.Print(message)
	}
}
