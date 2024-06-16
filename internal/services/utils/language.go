package utils

import (
	"fmt"
	"unicode"
)

func DetermineLanguage(str string) (unicodeRange *unicode.RangeTable, err error) {
	ruCount, engCount := 0, 0
	for _, r := range str {
		if unicode.Is(unicode.Cyrillic, r) {
			ruCount++
		} else if unicode.Is(unicode.Latin, r) {
			engCount++
		} else if !unicode.In(r, unicode.Cyrillic, unicode.Latin, unicode.Number, unicode.Space, unicode.Punct) {
			return nil, fmt.Errorf("the string contains an unsupported character: \"%s\"", string(r))
		}
	}
	if ruCount > engCount {
		unicodeRange = unicode.Cyrillic
	} else {
		unicodeRange = unicode.Latin
	}
	return unicodeRange, err
}
