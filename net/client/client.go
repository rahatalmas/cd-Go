package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func ReceiveMsg(conn net.Conn, ch chan string) {
	for {
		arr := make([]byte, 128)
		_, e := conn.Read(arr)
		if e != nil {
			fmt.Println(e)
		}
		fmt.Println(string(arr))
	}
}

type message struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

func main() {
	m := message{}
	fmt.Println("enter your name")
	fmt.Scanln(&m.User)
	m.Msg = "connect me"

	fmt.Println("client")
	ch := make(chan string)
	conn, err := net.Dial("tcp", "localhost:6000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println(conn.RemoteAddr().String())

	go ReceiveMsg(conn, ch)

	me, _ := json.Marshal(m)
	_, er := conn.Write(me)
	if er != nil {
		fmt.Println("Error sending message:", er)
	}

	for {
		scanner := bufio.NewScanner(os.Stdin)
		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if !scanner.Scan() {
			fmt.Println("Not able to scan the input")
		}
		str := scanner.Text()

		m.Msg = str
		jsonData, _ := json.Marshal(m)
		_, err := conn.Write(jsonData)
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}

		if str == "exit" {
			fmt.Println("Exiting...")
			break
		}
		fmt.Println(len(ch))
	}
}

//taking inputs:

//normal way
// var str string
// fmt.Scan(&str)

//scanner //can get data a full line ... until new line
// scanner := bufio.NewScanner(os.Stdin)
// if !scanner.Scan() {
// 	fmt.Println("Not able to scan the input")
// }
// if err := scanner.Err(); err != nil {
// 	fmt.Println("Error reading input:", err)
// 	return
// }
// message := scanner.Text()

//can get data a full line ... until new line
// reader := bufio.NewReader(os.Stdin)
// 	for {
// 		text, _ := reader.ReadString('\n')
// 		_, err := conn.Write([]byte(text))
// 		if err != nil {
// 			fmt.Println("Error sending message:", err)
// 			return
// 		}
// 		if text == "exit\n" {
// 			fmt.Println("Exiting...")
// 			break
// 		}
// 	}
