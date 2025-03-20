package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world")
	// b := bytes.NewReader([]byte("rahat almas")) // don't erase the data if i raed
	// fmt.Println(b)
	buff := bytes.Buffer{}
	//buff.Grow(1000)
	fmt.Println(buff.Cap())
	fmt.Println(buff.Len())

	buff.Write([]byte("rahat almas")) // erase the data if i read

	fmt.Println(buff.Cap())
	fmt.Println(buff.Len())

	a := make([]byte, 3)
	for {
		_, e := buff.Read(a)
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(string(bytes.ToUpper(a)))
	}
	s := strings.NewReader("hello from mars")
	fmt.Println(s)

}
