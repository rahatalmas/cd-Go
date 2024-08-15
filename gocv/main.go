package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	// Open a new window
	window := gocv.NewWindow("Hello World")
	defer window.Close()

	// Load an image from file
	img := gocv.IMRead("n.jpg", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Error reading image from file")
		return
	}
	defer img.Close()

	// Show image
	window.IMShow(img)
	window.WaitKey(0)
}
