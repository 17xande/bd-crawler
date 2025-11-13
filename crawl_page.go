package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

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

	pageData := extractPageData(html, rawCurrentURL)
	cfg.setPageData(normalized, pageData)

	for _, url := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}
