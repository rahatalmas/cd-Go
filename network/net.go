package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

var (
	clients = make(map[string]net.Conn) // Track connected clients                // Mutex to synchronize access to the clients map
)

func handleListener(conn net.Conn) {
	defer conn.Close()
	address := conn.RemoteAddr().String()

	// Register the client
	clients[address] = conn
	for k, v := range clients {
		fmt.Println(k, " ", v)
	}

	fmt.Println("Client connected:", address)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Sender:", address, "Message:", message)

		if message == "c" {
			f, err := os.Create("codes.txt")
			if err != nil {
				conn.Write([]byte("failed creating file"))
			}
			_, er := conn.Write([]byte(f.Name() + " created\n"))
			if er != nil {
				fmt.Println(er)
			}
		} else if message == "list" {
			for addr, _ := range clients {
				conn.Write([]byte(addr))
			}
		} else {
			_, err := clients[message].Write([]byte("hello hehe" + "\n"))
			if err != nil {
				fmt.Println("Response send error:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner error:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5050")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Server is listening on port 5050")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept error:", err)
			continue
		}
		go handleListener(conn)
	}
}
