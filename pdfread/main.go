package main

import (
    "fmt"
    "log"
    "github.com/pdfcpu/pdfcpu"
    "github.com/pdfcpu/pdfcpu/pkg/api"
    "github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
    "strings"
)

func main() {
    // Path to your PDF file
    pdfPath := "sample.pdf"

    // Read the PDF file
    ctx, err := api.ReadContextFile(pdfPath)
    if err != nil {
        log.Fatalf("error reading PDF file: %s\n", err)
    }

    // Extract text from the PDF
    text, err := api.ReadContext(pdfPath, nil)
    if err != nil {
        log.Fatalf("error extracting text from PDF file: %s\n", err)
    }

    // Print the raw text
    fmt.Println("Raw text extracted from the PDF:")
    fmt.Println(text)

    // Process the extracted text (simple row-based extraction)
    fmt.Println("\nText organized by rows:")
    rows := strings.Split(text, "\n")
    for i, row := range rows {
        fmt.Printf("Row %d: %s\n", i+1, row)
    }
}
