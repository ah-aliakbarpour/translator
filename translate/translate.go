package translate

type Result struct {
	Source       string
	Translations []string
}

type Translator interface {
	Translate() ([]Result, error)
}
