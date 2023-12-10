package bb

import (
	"strings"
	"unicode"
)

func StringContainsAnySubstring(s string, substrs []string) bool {
	for _, ss := range substrs {
		if strings.Contains(s, ss) {
			return true
		}
	}
	return false
}

func StringContainsNonPrintableCharacters(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return true
		}
	}
	return false
}

func RemoveNonPrintableCharactersFromString(s string) string {
	var result []rune
	for _, r := range s {
		if unicode.IsPrint(r) {
			result = append(result, r)
		}
	}
	s = string(result)
	return s
}
