package main

import (
	"fmt"
	"os"
)

func main() {
	// Ensure there's exactly one URL argument
	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	// Get the URL argument
	url := os.Args[1]
	pages := make(map[string]int)

	fmt.Println("starting crawl")

	// Get the HTML body
	page := crawlPage(url, url, pages, 10)

	fmt.Printf("\nURLs and their Counts:\n--------------------------------------------\n\n")
	// Print the extracted URLs
	for extractedURL, count := range page {
		fmt.Printf("Extracted URL: %s\nCount: %d\n\n", extractedURL, count)
	}
}
