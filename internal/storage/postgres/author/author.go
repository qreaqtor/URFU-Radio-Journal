package authorst

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
)

type AuthorStorage struct {
	db    *sql.DB
	table string
}

func NewAuthorStorage(db *sql.DB, table string) *AuthorStorage {
	return &AuthorStorage{
		db:    db,
		table: table,
	}
}

func (as *AuthorStorage) InsertMany(authors []models.Author, articleID string) error {
	for _, author := range authors {
		var authorID int

		query := fmt.Sprintf(
			"INSERT INTO %s (fullname_ru, fullname_en, affiliation, email, article_id) VALUES ($1, $2, $3, $4, $5) RETURNING author_id",
			as.table,
		)
		row := as.db.QueryRow(
			query,
			author.FullName.Ru,
			author.FullName.Eng,
			author.Affilation,
			author.Email,
			articleID,
		)

		err := row.Scan(&authorID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (as *AuthorStorage) Find(articleID string) ([]models.Author, error) {
	var authors []models.Author
	query := fmt.Sprintf(
		"SELECT fullname_ru, fullname_en, affiliation, email FROM %s WHERE article_id = $1",
		as.table,
	)
	rows, err := as.db.Query(query, articleID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author models.Author
		err := rows.Scan(
			&author.FullName.Ru,
			&author.FullName.Eng,
			&author.Affilation,
			&author.Email,
		)

		if err != nil {
			return nil, err
		}

		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (as *AuthorStorage) Delete(articleID string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE article_id = $1",
		as.table,
	)

	_, err := as.db.Exec(query, articleID)
	if err != nil {
		return err
	}

	return err
}
