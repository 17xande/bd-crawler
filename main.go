package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	const usage = "usage: crawler <url> <maxConcurrency> <maxPages>"
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(usage)
		os.Exit(1)
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	fmt.Printf("cfg: %#v\n", cfg)
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
