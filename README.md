# SearchCrawler

Ein einfaches Go-Tool, das Suchergebnisse über DuckDuckGo abruft, die verlinkten Seiten besucht und versucht, den Titel der jeweiligen Webseite auszulesen. Die Ergebnisse werden in einer JSON-Datei gespeichert.

## Funktionsweise

Das Programm:

1. sucht über DuckDuckGo nach einem Begriff (Text wird per URL-Query übergeben),
2. extrahiert die tatsächlichen Ziel-URLs (die DuckDuckGo über Redirects maskiert),
3. besucht jede Seite mit einem realistischen User-Agent,
4. liest den `<title>`-Tag der Seite (wenn möglich),
5. speichert URL + Titel als Liste in eine `results.json`.

Die Ergebnisse erscheinen in der Reihenfolge, in der sie besucht wurden.

## Verwendung

```bash
go run main.go "schuhe"
```

Das Skript erstellt oder überschreibt eine Datei namens `results.json` im Projektverzeichnis.

## Beispielausgabe

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

Wenn kein Titel extrahiert werden konnte, bleibt das Feld leer oder fehlt.

## Technisches

- Sprache: Go
- HTTP/Scraping: [Colly](https://github.com/gocolly/colly)
- JSON-Encoding über Standardbibliothek
- Einfache Fehlerbehandlung und Logging per Konsole

## Hinweise

Nicht alle Seiten erlauben Bot-Zugriffe. Manche blockieren Crawler oder liefern keine Inhalte ohne JavaScript. Das Tool ist nicht dafür gedacht, intensiv zu scrapen oder Daten großflächig zu sammeln – es dient eher als Lernprojekt oder für vereinfachte Recherchezwecke.

## Lizenz / Autor

Keine besondere Lizenz. Nur zu Lernzwecken. Erstellt von Furkan.