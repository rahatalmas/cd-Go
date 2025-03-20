package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Connecting to server...")
	conn, err := net.Dial("tcp", "localhost:5050")
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to", conn.RemoteAddr().String())
	fmt.Println("Enter your message:")

	go receiveMessages(conn) // Start receiving messages in a goroutine

	reader := bufio.NewReader(os.Stdin)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}
	}
}

func receiveMessages(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println("Server response:", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Read error:", err)
	}
}
