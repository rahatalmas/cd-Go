package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func err(e any) {
	if e != nil {
		fmt.Println(e)
	}
}

type message struct {
	User string `json:"user"`
	Msg  string `json:"msg"`
}

func ConnectionHandler(conn net.Conn) {
	//fmt.Println(conn.RemoteAddr().Network())

	for {
		arr := make([]byte, 1024)
		_, e := conn.Read(arr)
		if e != nil && e != io.EOF {
			break
		}
		fmt.Println(string(arr))
		conn.Write([]byte("hello from server"))
	}
}

func main() {
	fmt.Println("Hello World")
	users := make(map[string]net.Conn)

	s, e := net.Listen("tcp", ":6000")
	err(e)
	defer s.Close()
	for {
		conn, err := s.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(conn.RemoteAddr().String())

		arr := make([]byte, 1024)
		n, e := conn.Read(arr)
		if e != nil && e != io.EOF {
			break
		}
		//fmt.Println(string(arr))
		m := message{}
		der := json.Unmarshal(arr[0:n], &m)
		if der != nil {
			fmt.Println(der)
		}
		fmt.Println(m.User)
		users[m.User] = conn
		fmt.Println(users)

		conn.Write([]byte("hello from server"))
		ul, mer := json.Marshal(users)
		if mer != nil {
			fmt.Println(mer)
		}
		conn.Write(ul)
		go ConnectionHandler(conn)
	}
}

//need to learn for optimization or better control
// buff := bytes.Buffer{}
// buff.Grow(1)
//var buff bytes.Buffer
// _, er := io.Copy(&buff, conn)
// if er != nil {
// 	fmt.Println(e)
// }
// fmt.Println(string(arr))
// conn.Read(buff.Bytes())
// fmt.Println(buff.Len())
// fmt.Println(buff.String())
// for {
// 	_, e := buff.Read(arr)
// 	if e != nil {
// 		fmt.Println(e)
// 		break
// 	}
// 	fmt.Println(string(arr))
// }
