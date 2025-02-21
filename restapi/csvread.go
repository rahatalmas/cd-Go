package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func H() {
	file, err := os.Open("data.csv")
	fmt.Println(file.Name())
	if err != nil {
		fmt.Println("file open failed")
	}
	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}
