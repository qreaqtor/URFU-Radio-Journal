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

	generate := getGenerateFunc(language, key)

	if checkBoolPtr(args.Title) {
		query += generate("title")
	}

	if checkBoolPtr(args.Affilation) {
		query += generate("content")
	}

	// if checkBoolPtr(args.Authors) {
	// 	query += generate("jsonb_array_elements_text(authors->'fullname')")
	// }

	// if checkBoolPtr(args.Keywords) {
	// 	query += generate("jsonb_array_elements_text(keywords)")
	// }

	return query, true
}

func checkBoolPtr(b *bool) bool {
	return b != nil && *b
}

func getGenerateFunc(language, key string) func(field string) string {
	i := 0
	return func(field string) string {
		i++
		condition := fmt.Sprintf("to_tsvector('%s', %s->>'%s') @@ websearch_to_tsquery('%s', $1)", language, field, key, language)
		if i > 1 {
			return " OR " + condition
		}
		return condition
	}
}
