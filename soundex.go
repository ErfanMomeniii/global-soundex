package globalsoundex

func New(language string) Soundex {
	return soundexBasedOnLanguage(language)
}

type Soundex interface {
	Encode(string) string
	Suggest(string) string
	Correspond(string) string
	AddEntities([]string)
}
