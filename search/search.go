package search

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"net/url"
	"strings"
)

func convertRawToLink(rawLink string) (string, error) {
	parsed, err := url.Parse(rawLink)
	if err != nil {
		return "", err
	}
	uddg := parsed.Query().Get("uddg")
	if uddg == "" {
		return "", fmt.Errorf("kein uddg gefunden")
	}
	decoded, err := url.QueryUnescape(uddg)
	if err != nil {
		return "", err
	}
	return decoded, nil
}

func Search(searchTerm string) []string {
	var results []string
	c := colly.NewCollector()

	// Handler für Suchergebnisse
	c.OnHTML("a.result__a", func(e *colly.HTMLElement) {
		rawLink := e.Attr("href")
		cleanLink, _ := convertRawToLink(rawLink)
		if strings.Contains(cleanLink, "duckduckgo.com/y.js") ||
			strings.Contains(cleanLink, "click_metadata") ||
			strings.Contains(cleanLink, "bing.com/aclick") {
			return
		}
		results = append(results, cleanLink)
	})

	searchUrl := fmt.Sprintf("https://duckduckgo.com/html/?q=%s", searchTerm)
	if err := c.Visit(searchUrl); err != nil {
		fmt.Println("Fehler bei der Suche:", searchUrl, err)
	}
	c.Wait()

	return results
}

// TODO 1. bot detection Erkennung einbauen
// könnte sein, dass duckduckgo uns als bot erkennt und blockiert und deshalb fehler entstehen wie timeout oder 403
