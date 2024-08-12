package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func h() {
	file, err := os.Open("data.csv")
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
