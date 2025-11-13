package main

import (
	"fmt"
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
	const maxConcurrency = 3

	cfg, err := configure(baseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("error getting config: %v", err)
		os.Exit(1)
	}

	fmt.Printf("starting crawl of: %s\n", baseURL)
	cfg.wg.Add(1)
	cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	fmt.Println("\nPage: Count\n------------")
	for key, val := range cfg.pages {
		fmt.Printf("%s:\t%v\n", key, val)
	}
}
