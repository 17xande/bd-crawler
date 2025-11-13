package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
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

	if isFirst := cfg.addPageVisit(normalized); !isFirst {
		return
	}

	fmt.Printf("crawling %s\n", rawCurrentURL)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html: %v\n", err)
		return
	}
	cfg.mu.Lock()

	page := cfg.pages[normalized]

	URLs, err := getURLsFromHTML(html, currentURL)
	if err != nil {
		fmt.Printf("error getting URLs from html: %v\n", err)
		return
	}
	page.OutgoingLinks = URLs
	cfg.pages[normalized] = page
	cfg.mu.Unlock()

	for _, url := range URLs {
		cfg.wg.Add(1)
		defer func() {
			cfg.wg.Done()
			<-cfg.concurrencyControl
		}()

		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
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
