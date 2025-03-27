# 🔎 SearchCrawler

Ein einfacher Webcrawler in Go, der über DuckDuckGo nach einem Begriff sucht, die gefundenen Webseiten besucht, deren Titel extrahiert und alle Ergebnisse als JSON-Datei speichert.

---

## 📦 Features

- Suchbegriff über DuckDuckGo
- Besuch der gefundenen Webseiten
- Titel-Extraktion (HTML `<title>`)
- JSON-Ausgabe: `url` + optional `title`
- Ergebnisse werden in Besuchsreihenfolge gespeichert
- Bestehende `results.json` wird immer überschrieben

---

## 📁 Projektstruktur

```
.
├── main.go           # Einstiegspunkt, ruft Suche & Crawler auf
├── search.go         # DuckDuckGo-Suche + Link-Parsing
├── crawler.go        # Besucht Seiten & extrahiert Titel
├── results.json      # JSON-Ausgabe (automatisch erstellt)
├── go.mod / go.sum   # Abhängigkeiten (Colly etc.)
```

---

## ▶️ Verwendung

```bash
go run main.go "schuhe"
```

Das Programm sucht bei DuckDuckGo nach dem Begriff `schuhe`, besucht jede relevante Seite und speichert die Ergebnisse in `results.json`.

---

## 📄 Beispielausgabe (`results.json`)

```json
[
  {
    "url": "https://www.otto.de/schuhe/",
    "title": "Schuhe online kaufen | OTTO"
  },
  {
    "url": "https://www.eschuhe.de/"
  }
]
```

Hinweis: Wenn der Titel nicht extrahiert werden kann, erscheint nur die URL.

---

## 🔧 Nächste Schritte (To Do)

- [ ] Webhook-Benachrichtigung (Discord, GitHub Webhooks)
- [ ] Fehler besser handhaben & anzeigen
- [ ] Bot-Tarnung (User-Agent, Verzögerungen)

---

## 📚 Technologien

- [Go](https://golang.org/)
- [Colly](https://github.com/gocolly/colly) – Web Scraping Framework

---

## 🧠 Gelernt beim Projekt

- Web-Scraping mit Colly
- JSON-Serialisierung in Go
- Umgang mit Structs, Maps, Slices und Pointern
- URL-Encoding & DuckDuckGo-Redirects verstehen

---

## 👨‍💻 Autor

Furkan – Projekt zum Lernen von Go & Webscraping  
