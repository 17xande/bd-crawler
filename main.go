package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Printf("error parsing url: %v", err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]PageData),
		baseURL:            parsedURL,
		concurrencyControl: make(chan struct{}),
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)
	cfg.crawlPage(baseURL)

	fmt.Println("\nPage: Count\n------------")
	for key, val := range cfg.pages {
		fmt.Printf("%s:\t%v\n", key, val)
	}
}
