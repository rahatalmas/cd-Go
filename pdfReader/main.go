package main

import (
    "fmt"
    "log"

    "github.com/pdfium/pdfium"
)

func main() {
    // Path to the PDF file
    pdfFile := "example.pdf"

    // Initialize PDFium
    pdfium.Init()
    defer pdfium.Destroy()

    // Load the PDF document
    doc, err := pdfium.LoadDocument(pdfFile)
    if err != nil {
        log.Fatalf("error loading PDF document: %s\n", err)
    }
    defer doc.Close()

    // Get the number of pages
    pageCount := doc.GetPageCount()

    // Iterate through the pages and extract text
    for i := 0; i < pageCount; i++ {
        page := doc.GetPage(i)
        if page == nil {
            log.Printf("error getting page %d\n", i)
            continue
        }

        // Extract text from the page
        text, err := page.GetText()
        if err != nil {
            log.Printf("error extracting text from page %d: %s\n", i, err)
            continue
        }

        fmt.Printf("Text from page %d:\n%s\n", i+1, text)
    }
}

