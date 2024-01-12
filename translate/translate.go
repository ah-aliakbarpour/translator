package translate

type Word struct {
	Source     string
	Translates []string
}

func Translate(sourceLanguage, translateLanguage string, words []string) ([]Word, error) {
	return []Word{}, nil
}
