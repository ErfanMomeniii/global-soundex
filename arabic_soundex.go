package globalsoundex

import (
	"math"
	"regexp"
	"strings"
	"unicode"
)

type ArabicSoundex struct {
	mappedDatas map[string]string
}

const tokenLengthAr = 4

var (
	arabicMapping = map[rune]string{
		'ب': "1", 'ف': "1",
		'خ': "2", 'ج': "2", 'ز': "2", 'س': "2", 'ش': "2", 'ص': "2", 'غ': "2", 'ق': "2", 'ک': "2",
		'ت': "3", 'ث': "3", 'د': "3", 'ذ': "3", 'ض': "3", 'ط': "3", 'ظ': "3",
		'ل': "4",
		'م': "5", 'ن': "5",
		'ر': "6",
	}

	translit = map[rune]string{
		'ا': "A", 'ب': "B", 'ت': "T", 'ث': "T", 'ج': "J", 'ح': "H", 'خ': "K",
		'د': "D", 'ذ': "Z", 'ر': "R", 'ز': "Z", 'س': "S", 'ش': "S", 'ص': "S",
		'ض': "D", 'ط': "T", 'ظ': "Z", 'ع': "A", 'غ': "G", 'ف': "F", 'ق': "Q",
		'ك': "K", 'ل': "L", 'م': "M", 'ن': "N", 'ه': "H", 'و': "W", 'ي': "Y", 'ء': "A",
	}

	replacements = map[string]string{
		"آ": "ا", "أ": "ا", "إ": "ا",
		"ى": "ي", "ة": "ت",
	}

	commonPrefixes = []string{"ال", "ابو", "بن", "عبد", "ام", "ابن"}
)

func (s *ArabicSoundex) normalize(inputStr string) string {
	for old, newVal := range replacements {
		inputStr = strings.ReplaceAll(inputStr, old, newVal)
	}

	diacritics := regexp.MustCompile("[\u064B-\u0652\u0670\u06D6-\u06ED]")
	inputStr = diacritics.ReplaceAllString(inputStr, "")
	inputStr = strings.ReplaceAll(inputStr, "ـ", "") // Tatweel

	return inputStr
}

func (s *ArabicSoundex) stripCommonPrefixes(inputStr string) string {
	for _, prefix := range commonPrefixes {
		if strings.HasPrefix(inputStr, prefix) {
			return strings.TrimPrefix(inputStr, prefix)
		}
	}
	return inputStr
}

func (s *ArabicSoundex) filterArabic(inputStr string) string {
	var result strings.Builder
	for _, char := range inputStr {
		if unicode.Is(unicode.Arabic, char) {
			result.WriteRune(char)
		}
	}
	return result.String()
}

func (s *ArabicSoundex) removeDuplicate(inputStr string) string {
	var output strings.Builder
	var lastCode string
	first := true

	for _, ch := range inputStr {
		code, exists := arabicMapping[ch]
		if !exists {
			continue
		}
		if first {
			lastCode = code
			output.WriteRune(ch)
			first = false
			continue
		}
		if code != lastCode {
			output.WriteRune(ch)
			lastCode = code
		}
	}
	return output.String()
}

func (s *ArabicSoundex) mask(inputStr string) string {
	var output strings.Builder
	for _, c := range inputStr {
		if val, ok := arabicMapping[c]; ok {
			output.WriteString(val)
		}
	}

	for output.Len() < tokenLengthAr-1 {
		output.WriteByte('0')
	}
	return output.String()[:tokenLengthAr-1]
}

func (s *ArabicSoundex) Encode(inputStr string) string {
	inputStr = s.filterArabic(inputStr)

	inputStr = s.stripCommonPrefixes(s.normalize(inputStr))
	runes := []rune(inputStr)
	if len(runes) == 0 {
		return ""
	}

	return translit[runes[0]] + s.mask(s.removeDuplicate(string(runes[1:])))
}

func (s *ArabicSoundex) AddEntities(entities []string) {
	for _, entity := range entities {
		code := s.Encode(entity)
		if _, ok := s.mappedDatas[code]; !ok {
			s.mappedDatas[code] = entity
		}
	}
}

func (s *ArabicSoundex) Correspond(inputStr string) string {
	if m, ok := s.mappedDatas[s.Encode(inputStr)]; ok {
		return m
	}
	return inputStr
}

func (s *ArabicSoundex) Suggest(inputStr string) string {
	if inputStr == "" {
		return ""
	}

	output := inputStr
	e := s.Encode(inputStr)
	distance := tokenLengthAr + 1

	if len(s.mappedDatas) == 0 {
		return output
	}

	for md, entity := range s.mappedDatas {
		if e == md {
			return entity
		}
		d := levenshtein(e, md)
		if d < distance {
			distance = d
			output = entity
		}
	}

	return output
}

func findMin(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}

func levenshtein(s1, s2 string) int {
	lenS1 := len(s1)
	lenS2 := len(s2)

	dp := make([][]int, lenS1+1)
	for i := range dp {
		dp[i] = make([]int, lenS2+1)
	}

	for i := 0; i <= lenS1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= lenS2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= lenS1; i++ {
		for j := 1; j <= lenS2; j++ {
			cost := 0
			if s1[i-1] != s2[j-1] {
				cost = 1
			}
			dp[i][j] = findMin(
				dp[i-1][j]+1,
				dp[i][j-1]+1,
				dp[i-1][j-1]+cost,
			)
		}
	}

	return dp[lenS1][lenS2]
}

func newArabicSoundex() *ArabicSoundex {
	return &ArabicSoundex{
		mappedDatas: make(map[string]string),
	}
}
