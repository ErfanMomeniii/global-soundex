package globalsoundex

import (
	"regexp"
	"strings"
)

type EnglishSoundex struct {
	mappedDatas map[string]string
}

const tokenLengthEn = 4

var (
	englishMapping = map[rune]string{
		'b': "1", 'f': "1", 'p': "1", 'v': "1",
		'c': "2", 'g': "2", 'j': "2", 'k': "2", 'q': "2", 's': "2", 'x': "2", 'z': "2",
		'd': "3", 't': "3",
		'l': "4",
		'm': "5", 'n': "5",
		'r': "6",
	}

	ignoredChars = "aeiouhwy"
)

func (s *EnglishSoundex) removeIgnoredChars(inputStr string) string {
	var output string

	for _, c := range inputStr {
		if !strings.Contains(ignoredChars, string(c)) {
			output += string(c)
		}
	}

	return output
}

func (s *EnglishSoundex) removeDuplicate(inputStr string) string {
	var lastSee, output string

	lastSee = englishMapping[rune(inputStr[0])]
	output += string(inputStr[0])

	for i := 1; i < len(inputStr); i++ {
		if lastSee != englishMapping[rune(inputStr[i])] {
			lastSee = englishMapping[rune(inputStr[i])]
			output += string(inputStr[i])
		}
	}

	return output
}

func (s *EnglishSoundex) mask(inputStr string) string {
	var output string

	for _, c := range inputStr {
		output += englishMapping[c]
	}

	for i := len(output); i < tokenLengthEn-1; i++ {
		output += "0"
	}

	return output[:tokenLengthEn-1]
}

func (s *EnglishSoundex) Encode(inputStr string) string {
	inputStr = strings.ToLower(regexp.MustCompile("[^a-zA-Z]").ReplaceAllString(strings.TrimSpace(inputStr), ""))
	if inputStr == "" {
		return inputStr
	}

	return string(inputStr[0]) + s.mask(s.removeDuplicate(s.removeIgnoredChars(inputStr[1:])))
}

func (s *EnglishSoundex) AddEntities(entities []string) {
	for _, entity := range entities {
		if _, ok := s.mappedDatas[s.Encode(entity)]; !ok {
			s.mappedDatas[s.Encode(entity)] = entity
		}
	}
}

func (s *EnglishSoundex) Correspond(inputStr string) string {
	if m, ok := s.mappedDatas[s.Encode(inputStr)]; ok {
		return m
	}

	return inputStr
}

func hammingDistance(s1, s2 string) int {
	var distance int

	for i := 0; i < tokenLengthEn; i++ {
		if s1[i] != s2[i] {
			distance++
		}
	}

	return distance
}

func (s *EnglishSoundex) Suggest(inputStr string) string {
	output := inputStr
	e := s.Encode(inputStr)
	distance := tokenLengthEn + 1

	for md := range s.mappedDatas {
		d := hammingDistance(e, md)
		if d < distance {
			distance = d
			output = s.mappedDatas[md]
		}
	}

	return output
}

func newEnglishSoundex() *EnglishSoundex {
	return &EnglishSoundex{
		mappedDatas: make(map[string]string),
	}
}
