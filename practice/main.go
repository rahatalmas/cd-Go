package main

import (
	"fmt"
)

func odd(ch chan int) {
	for i := 1; i < 100; i += 2 {
		ch <- i
	}
	close(ch)
}

func even(ch chan int) {
	ch <- 2
}

func message(ch chan int) {
	for {
		fmt.Println("message 1")
		var m int
		fmt.Scanln(&m)
		ch <- m
	}
}

func message2(ch chan int) {
	for {
		fmt.Println("message 2")
		var m int
		fmt.Scanln(&m)
		ch <- m
	}
}

func main() {
	fmt.Println("hello pretty <3")
	ch := make(chan int)
	go message(ch)
	go message2(ch)
	//go odd(ch)
	//go even(ch)
	// var result int
	// result = <-ch
	// fmt.Println(result)
	// result = <-ch
	// fmt.Println(result)
	for i := range ch {
		fmt.Println("from go routine ", i)
	}
}
