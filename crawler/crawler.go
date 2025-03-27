package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
	"strings"
)

type PageResult struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func Crawl(urls []string) {
	c := colly.NewCollector()
	var results []PageResult
	currentUrl := ""

	c.OnHTML("title", func(e *colly.HTMLElement) {
		title := strings.TrimSpace(e.Text)
		if title == "" {
			title = "Kein Titel gefunden"
		}
		results = append(results, PageResult{
			Title: title,
			URL:   currentUrl,
		})
		fmt.Printf("Titel: %s\n", title)
		fmt.Println("────────────────────────────────────────")
	})

	c.OnRequest(func(r *colly.Request) {
		currentUrl = r.URL.String()
		fmt.Println("Besuche:", currentUrl)
	})

	for _, rawUrl := range urls {
		if err := c.Visit(rawUrl); err != nil {
			fmt.Println("Fehler beim Besuch:", rawUrl)
		}
	}

	c.Wait()

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
