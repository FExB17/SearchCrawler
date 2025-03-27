package main

import (
	"SearchCrawler/crawler"
	"SearchCrawler/search"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Verwende: go run main.go [Suchbegriff]")
		return
	}
	searchTerm := os.Args[1]
	results := search.Search(searchTerm)
	crawler.Crawl(results)
}
