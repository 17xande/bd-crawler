package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		panic(err)
	}

	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		panic(err)
	}

	fmt.Printf("adding page visit for %s\n", normalized)
	if isFirst := cfg.addPageVisit(normalized); !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v\n", err)
		return
	}

	page := cfg.pages[normalized]

	URLs, err := getURLsFromHTML(html, currentURL)
	if err != nil {
		fmt.Printf("error getting URLs from html: %v\n", err)
		return
	}
	page.OutgoingLinks = URLs
	cfg.pages[normalized] = page

	for _, url := range URLs {
		cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	_, found := cfg.pages[normalizedURL]
	isFirst = !found
	if found {
		return
	}
	page := PageData{
		URL: normalizedURL,
	}
	cfg.pages[normalizedURL] = page
	return
}
