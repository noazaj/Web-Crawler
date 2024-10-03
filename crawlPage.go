package main

import (
	"fmt"
	"net/url"
	"time"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int, depth int) map[string]int {
	if depth == 0 {
		return pages
	}

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("error parsing base URL during crawl: %v", err)
	}

	current, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("error parsing current URL during crawl: %v", err)
	}

	if base.Host != current.Host {
		return pages
	}

	normalizedCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error normalizing current URL: %v", err)
	}

	_, ok := pages[normalizedCurrentURL]
	if ok {
		pages[normalizedCurrentURL]++
		return pages
	} else {
		pages[normalizedCurrentURL] = 1
	}

	// Print what URL we are crawling and sleep for some time
	// to not overwork the server
	fmt.Printf("Crawling %s\n", normalizedCurrentURL)

	time.Sleep(500 * time.Millisecond)

	currentHTML, err := getHTML(normalizedCurrentURL)
	if err != nil {
		fmt.Printf("error getting HTML for current URL: %v", err)
	}

	allURLs, err := getURLsFromHTML(currentHTML, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting all URLs from HTML of current URL: %v", err)
	}

	for _, url := range allURLs {
		crawlPage(rawBaseURL, url, pages, depth-1)
	}

	return pages
}
