package main

import (
	"fmt"
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
}

func extractPageData(html, pageURL string) PageData {
	data := PageData{
		URL:            pageURL,
		H1:             getH1FromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
	}

	parsedURL, err := url.Parse(pageURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return data
	}
	links, err := getURLsFromHTML(html, parsedURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return data
	}
	data.OutgoingLinks = links
	imageURLs, err := getImagesFromHTML(html, parsedURL)
	if err != nil {
		fmt.Printf("%v\n", err)
		return data
	}
	data.ImageURLs = imageURLs
	return data
}
