package main

import (
	"fmt"
	"net/url"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v%v", parsedURL.Host, parsedURL.Path), nil
}