package main

import (
	"fmt"
)

func Input(ch chan string) {
	for {
		var s string
		fmt.Scan(&s)
		ch <- s
	}
}

func Output(ch chan string) {
	for {
		m := <-ch
		fmt.Println(m)
	}
}

func main() {
	fmt.Println("hello world")
	ch := make(chan string)
	// go Input(ch)
	// go Output(ch)
	//select {}
	// for {
	// 	m := <-ch
	// 	fmt.Println(m)
	// }
	for {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			var s string
			fmt.Scan(&s)
			ch <- s
		}
	}
}
