package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
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
		mu:                 &sync.Mutex{},
		pages:              make(map[string]PageData),
		baseURL:            parsedURL,
		concurrencyControl: make(chan struct{}, 5),
		wg:                 &sync.WaitGroup{},
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	fmt.Println("\nPage: Count\n------------")
	for key, val := range cfg.pages {
		fmt.Printf("%s:\t%v\n", key, val)
	}
}
