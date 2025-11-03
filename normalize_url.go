package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	fullpath := parsedURL.Host + parsedURL.Path
	fullpath = strings.ToLower(fullpath)
	fullpath = strings.TrimSuffix(fullpath, "/")
	return fullpath, nil
}
