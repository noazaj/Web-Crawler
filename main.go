package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello World")
	body := `<html>
						<body>
						<a href="/path/one">
							<span>Boot.dev</span>
						</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
						</body>
					</html>`
	urls, err := GetURLsFromHTML(body, "https://blog.boot.dev")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(urls)
}