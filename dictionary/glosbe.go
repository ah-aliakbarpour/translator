package dictionary

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"time"
)

type Glosbe struct {
	SourceLanguage, TranslateLanguage string
}

// Translate scrapes translations from Glosbe dictionary
func (dictionary *Glosbe) Translate(sourceWords []string) ([]Result, error) {
	collector := colly.NewCollector()

	results, err := dictionary.scrapeResults(collector, sourceWords)
	if err != nil {
		return nil, err
	}

	return results, nil
}

// scrapeResults iterates through SourceWords and stores the translation in result array
func (dictionary *Glosbe) scrapeResults(collector *colly.Collector, sourceWords []string) ([]Result, error) {
	const Domain = "glosbe.com"
	const TranslatedElementClass = "translation__item__pharse"

	var results []Result
	total := len(sourceWords)

	for i := 0; i < total; i++ {
		sourceWord := strings.ToLower(strings.TrimSpace(sourceWords[i]))
		if len(sourceWord) == 0 {
			continue
		}

		result := Result{Source: sourceWord}

		// find the translated elements
		collector.OnHTML("."+TranslatedElementClass, func(e *colly.HTMLElement) {
			result.Translations = append(result.Translations, strings.TrimSpace(e.Text))
		})

		// set request timeout
		collector.SetRequestTimeout(time.Minute)

		// visit the target page
		url := "https://" + Domain +
			"/" + dictionary.SourceLanguage +
			"/" + dictionary.TranslateLanguage +
			"/" + sourceWord
		err := collector.Visit(url)

		if errors.Is(err, colly.ErrAlreadyVisited) {
			// duplicate sourceWord
			continue
		} else if err != nil {
			if err.Error() != "Not Found" {
				return nil, fmt.Errorf("can't visit the target page colly: %w", err)
			}
		}

		results = append(results, result)
		printStatus(float64(i+1)/float64(total)*100, result)
	}

	return results, nil
}
