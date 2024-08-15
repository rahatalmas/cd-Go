package main

import (
	"dsa/tree"
	"fmt"
)

func main() {
	fmt.Println("hello pretty <3")
	t := tree.Create()
	n := 10
	var val int
	for i := 0; i < n; i++ {
		fmt.Scan(&val)
		t.Insert(val)
	}
	t.InOrderTraverse()
}
