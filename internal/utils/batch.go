package utils

import (
	"fmt"
	"urfu-radio-journal/internal/models"
)

// This function add limit and offset to query if args != 0
func AddBatchToQuery(query string, args *models.BatchArgs) string {
	if args.Offset > 0 {
		query += fmt.Sprintf(" OFFSET %v", args.Offset)
	}

	if args.Limit > 0 {
		query += fmt.Sprintf(" LIMIT %v", args.Limit)
	}

	return query
}
