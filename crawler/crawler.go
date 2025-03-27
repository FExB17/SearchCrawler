package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

type PageResult struct {
	Title string `json:"title,omitempty"`
	URL   string `json:"url"`
}

func Crawl(urls []string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/120.0.0.0 Safari/537.36"),
	)

	results := make(map[string]*PageResult)

	c.OnRequest(func(r *colly.Request) {
		url := r.URL.String()
		fmt.Println("Besuche:", url)

		if _, exists := results[url]; !exists {
			results[url] = &PageResult{
				URL: url,
			}
		}
	})

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		if title == "" {
			return
		}
		url := e.Request.URL.String()
		if entry, exists := results[url]; exists {
			entry.Title = title
		}
		fmt.Printf("Titel: %s\n", title)
		fmt.Println("────────────────────────────────────────")
	})

	for _, rawUrl := range urls {
		if err := c.Visit(rawUrl); err != nil {
			fmt.Println("Fehler beim Besuch:", rawUrl)
			fmt.Println("────────────────────────────────────────")
		}
	}

	c.Wait()

	final := make([]PageResult, 0, len(results))
	for _, entry := range results {
		final = append(final, *entry)
	}

	file, err := os.Create("results.json")
	if err != nil {
		fmt.Println("Fehler beim Erstellen der JSON-Datei:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Println("Fehler beim Schreiben in JSON:", err)
	}
}
