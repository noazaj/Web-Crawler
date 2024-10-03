package main

import (
	"io"
	"log"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		log.Fatalf("error making 'GET' request to URL: %s", rawURL)
		return "", err
	}
	defer resp.Body.Close()

	rawHTML, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading response body")
		return "", err
	}

	return string(rawHTML), nil
}
