package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(inputBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(inputBody))
	if err != nil {
		return nil, fmt.Errorf("can't create document from inputbody: %w", err)
	}
	links := []string{}
	_ = doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok {
			return
		}
		u, err := url.Parse(href)
		if err != nil {
			return
		}
		if u.Host == "" {
			u.Host = baseURL.Host
		}
		if u.Scheme == "" {
			u.Scheme = baseURL.Scheme
		}

		links = append(links, u.String())
	})
	return links, nil
}

func getImagesFromHTML(inputBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(inputBody))
	if err != nil {
		return nil, fmt.Errorf("can't create document from inputBody: %w", err)
	}
	links := []string{}
	_ = doc.Find("img[src]").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok {
			return
		}
		u, err := url.Parse(src)
		if err != nil {
			return
		}
		if u.Host == "" {
			u.Host = baseURL.Host
		}
		if u.Scheme == "" {
			u.Scheme = baseURL.Scheme
		}

		links = append(links, u.String())
	})
	return links, nil
}
