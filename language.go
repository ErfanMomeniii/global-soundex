package globalsoundex

import "strings"

const (
	Arabic  = "ar"
	English = "en" // default language
)

func parseLanguage(language string) string {
	if len(language) < 2 {
		return English
	}

	switch strings.ToLower(language[:2]) {
	case "ar":
		return Arabic
	default:
		return English
	}
}

func soundexBasedOnLanguage(language string) Soundex {
	lang := parseLanguage(language)

	switch lang {
	case "ar":
		return newArabicSoundex()
	default:
		return newEnglishSoundex()
	}
}
