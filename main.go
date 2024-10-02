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

	fmt.Println("starting crawl")
	fmt.Println(url)

	// Get the HTML body
	rawHTML, err := GetHTML(url)
	if err != nil {
		fmt.Printf("error getting HTMl body: %v", err)
		os.Exit(1)
	}

	// Parse the URLs from the HTML body
	urls, err := GetURLsFromHTML(rawHTML, url)
	if err != nil {
		fmt.Printf("Error parsing URLs from HTML: %v\n", err)
		os.Exit(1)
	}

	// Print the extracted URLs
	for _, extractedURL := range urls {
		fmt.Println(extractedURL)
	}
}
