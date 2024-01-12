package translate

type Translator struct {
	Sl, Tl      string
	SourceWords []string
}

type Result map[string][]string

func (translator *Translator) Translate() (Result, error) {
	return translator.googleTranslate()
}
