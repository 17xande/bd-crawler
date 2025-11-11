package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		panic(err)
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		panic(err)
	}

	if baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		panic(err)
	}
	if _, ok := pages[normalized]; ok {
		pages[normalized]++
		return
	}

	pages[normalized] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v\n", err)
		return
	}

	URLs, err := getURLsFromHTML(html, baseURL)
	if err != nil {
		fmt.Printf("error getting URLs from html: %v\n", err)
		return
	}

	for _, url := range URLs {
		crawlPage(rawBaseURL, url, pages)
	}
}
