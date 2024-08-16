package main

import (
	"dsa/tree"
	"fmt"
)

func main() {
	fmt.Println("hello pretty <3")
	for {
		tree.T()
		t := tree.Create()
		var n, val int
		fmt.Scanln(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&val)
			t.Insert(val)
		}
		t.InOrderTraverse()
		height := t.TreeHeight()
		fmt.Println("Depth of this Tree is: ", height)
		var u, v int
		fmt.Scanln(&u, &v)
		t.UpdateNode(u, v)
		t.InOrderTraverse()
	}
}
