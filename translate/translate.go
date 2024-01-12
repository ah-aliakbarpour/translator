package translate

type Result map[string][]string

type Translator interface {
	Translate() (Result, error)
}
