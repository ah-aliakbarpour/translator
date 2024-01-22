package dictionary

import "fmt"

type Result struct {
	Source       string
	Translations []string
}

type Dictionary interface {
	Translate(sourceWords []string) ([]Result, error)
}

func printStatus(percent float64, result Result) {
	fmt.Printf("\n[%.0f%%] %v: %v\n\n", percent, result.Source, result.Translations)
}
