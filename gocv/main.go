package main

import (
	"fmt"

	"gocv.io/x/gocv"
)

func main() {
	// OpenCV version
	fmt.Println("OpenCV Version:", gocv.Version())

	// Load an image from file
	img := gocv.IMRead("n.jpg", gocv.IMReadColor)
	if img.Empty() {
		fmt.Println("Error reading image")
		return
	}
	fmt.Println("Image loaded successfully")
}
