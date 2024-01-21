package dictionary

type Result struct {
	Source       string
	Translations []string
}

type Dictionary interface {
	Translate() ([]Result, error)
}
