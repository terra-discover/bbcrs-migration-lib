package lib

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// TrimSpace is exact same as strings.TrimSpace
func TrimSpace(str string) string {
	return strings.TrimSpace(str)
}

// TrimSpacePtr will trim the pointer string of space
func TrimSpacePtr(str *string) *string {
	if str != nil {
		trimmedStr := strings.TrimSpace(*str)
		return &trimmedStr
	}
	return nil
}

// TitleCase
// example : Pieces
func TitleCase(s string) string {
	t := cases.Title(language.Und)
	return t.String(s)
}

// SentenceCase
// example : 2 pieces up to 23 kg each
func SentenceCase(s string) string {
	firstWord := strings.Fields(s)[0]
	restOfSentence := strings.ToLower(s[len(firstWord):])
	return TitleCase(firstWord) + restOfSentence
}
