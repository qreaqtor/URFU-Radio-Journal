package utils

import (
	"fmt"
	"unicode"
	"urfu-radio-journal/internal/models"
)

// if args contains editionID then other fields will be ignored, because
// other fields should be used only for searching and editionID should be used only for getting edition's articles
func AddSearchToQuery(query string, args models.ArticleSearch) (string, bool) {
	query += " WHERE "

	if args.EditionID != 0 {
		return query + "edition_id = $1", false
	}

	unicodeRange := DetermineLanguage(args.Search)

	key := "Ru"
	language := "russian"
	if unicodeRange == unicode.Latin {
		key = "Eng"
		language = "english"
	}

	generate := getGenerateFunc(language)

	if checkBoolPtr(args.Title) {
		query += generate(fmt.Sprintf("title ->> '%s'", key))
	}

	if checkBoolPtr(args.Affilation) {
		query += generate(fmt.Sprintf("content ->>  '%s'", key))
	}

	if checkBoolPtr(args.Authors) {
		query += generate("authors")
	}

	if checkBoolPtr(args.Keywords) {
		query += generate("keywords")
	}

	return query, true
}

func checkBoolPtr(b *bool) bool {
	return b != nil && *b
}

func getGenerateFunc(language string) func(field string) string {
	i := 0
	return func(field string) string {
		i++
		condition := fmt.Sprintf("to_tsvector('%s', %s) @@ websearch_to_tsquery('%s', $1)", language, field, language)
		if i > 1 {
			return " OR " + condition
		}
		return condition
	}
}
