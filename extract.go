package main

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getH1FromHTML(html string) string {
	res, err := query(html, "h1")
	if err != nil {
		panic(err)
	}
	return res
}

func getFirstParagraphFromHTML(html string) string {
	res, err := query(html, "p")
	if err != nil {
		panic(err)
	}
	return res
}

func query(html, query string) (string, error) {
	reader := strings.NewReader(html)
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return "", fmt.Errorf("can't create document from reader: %w", err)
	}
	selection := document.Find(query).First()
	return selection.Text(), nil
}
