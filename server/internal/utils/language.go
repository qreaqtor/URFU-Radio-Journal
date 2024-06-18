package utils

import (
	"unicode"
)

func DetermineLanguage(str string) *unicode.RangeTable {
	ruCount, engCount := 0, 0
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			ruCount++
		} else if unicode.Is(unicode.Latin, r) {
			engCount++
		}
	}
	if ruCount > engCount {
		return unicode.Cyrillic
	}

	return unicode.Latin
}
