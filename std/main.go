package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("hello world")
	file,err := os.Open("std/data.txt")
	if err != nil{
		fmt.Println(err)
	}
	data := make([]byte,15)
	c,e := file.Read(data)
	if e!=nil{
		fmt.Println(e)
	}
	fmt.Println(c)
	fmt.Println(string(data))
}