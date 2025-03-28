package crawler

import (
	"SearchCrawler/tools"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

type PageResult struct {
	Title        string `json:"title,omitempty"`
	URL          string `json:"url"`
	BotDetection bool   `json:"bot_detection"`
}

func Crawl(urls []string) {
	c := colly.NewCollector(
		// User-Agent setzen, um wie ein Browser zu erscheinen
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
			"AppleWebKit/537.36 (KHTML, like Gecko) " +
			"Chrome/120.0.0.0 Safari/537.36"),
	)

	results := make(map[string]*PageResult)

	c.OnRequest(func(r *colly.Request) {
		url := r.URL.String()
		fmt.Println("Besuche:", url)
		// Header setzen, um wie ein Browser zu erscheinen
		r.Headers.Set("Accept-Language", "de-DE,de;q=0.9")
		r.Headers.Set("Referer", "https://www.google.com/")
		r.Headers.Set("DNT", "1")

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
		if entry, exists := results[url]; exists && entry.Title == "" {
			entry.Title = title
		}
		fmt.Printf("Titel: %s\n", title)
		fmt.Println("────────────────────────────────────────")
	})

	for _, rawUrl := range urls {
		detected := tools.CheckRobots(rawUrl)
		if _, exists := results[rawUrl]; !exists {
			results[rawUrl] = &PageResult{
				URL:          rawUrl,
				BotDetection: detected,
			}
		} else {
			results[rawUrl].BotDetection = detected
		}
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
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Fehler beim Schließen der JSON-Datei:", err)
		}
	}()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(results); err != nil {
		fmt.Println("Fehler beim Schreiben in JSON:", err)
	}
}
