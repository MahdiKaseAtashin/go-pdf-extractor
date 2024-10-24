package main

import (
	"fmt"
	"log"
	"os"

	pdf "github.com/ledongthuc/pdf"
)

func main() {
	// Open the PDF file
	file, err := os.Open("sample2.pdf")
	if err != nil {
		log.Fatalf("Error opening PDF file: %v", err)
	}
	defer file.Close()

	// Get file statistics to retrieve the file size
	fileStat, err := file.Stat()
	if err != nil {
		log.Fatalf("Error getting file information: %v", err)
	}

	// Create a new PDF reader
	reader, err := pdf.NewReader(file, fileStat.Size())
	if err != nil {
		log.Fatalf("Error creating PDF reader: %v", err)
	}

	var text string
	// Iterate through the PDF pages and extract text
	for i := 1; i <= reader.NumPage(); i++ {
		page := reader.Page(i)

		// Create an empty font map, as required by GetPlainText
		fonts := make(map[string]*pdf.Font)

		// Extract plain text from the page
		pageText, err := page.GetPlainText(fonts)
		if err != nil {
			log.Fatalf("Error creating PDF reader: %v", err)
			return
		}
		reorderedText, err := bidi.ProcessText(content, bidi.DefaultDirection)
		if err != nil {
			log.Printf("Error reordering text for page %d: %v", i, err)
			continue
		}

		persianText := "سلام دنیا" // Persian text meaning "Hello, world"
		fmt.Println(persianText)

		text += pageText

		// fmt.Printf("here: %v\n", pageText)
	}

	fmt.Printf("%v", text)

	// Print the extracted text
	// fmt.Println(text)
}
