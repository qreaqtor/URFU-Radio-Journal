package articlest

import (
	"database/sql"
	"fmt"
	"strconv"
	"urfu-radio-journal/internal/models"

	"github.com/lib/pq"
)

type ArticleStorage struct {
	db    *sql.DB
	table string
}

func NewArticleStorage(db *sql.DB, table string) *ArticleStorage {
	return &ArticleStorage{
		db:    db,
		table: table,
	}
}

func getColumns() string {
	return "edition_id, title_ru, title_en, content_ru, content_en, keywords_ru, keywords_en, file_path, video_path, literature, reference_ru, reference_en, date_receipt, date_acceptance, doi"
}

func generateValuesID(count int) string {
	res := ""
	for i := 1; i < count; i++ {
		res += fmt.Sprintf("$%v, ", i)
	}
	res += fmt.Sprintf("$%v", count)
	return res
}

func (as *ArticleStorage) InsertOne(article *models.ArticleCreate) (string, error) {
	var keywordsRu []string
	var keywordsEn []string

	for _, keyword := range article.Keywords {
		keywordsRu = append(keywordsRu, keyword.Ru)
		keywordsEn = append(keywordsEn, keyword.Eng)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING article_id;",
		as.table,
		getColumns(),
		generateValuesID(15),
	)
	row := as.db.QueryRow(
		query,
		article.EditionId,
		article.Title.Ru,
		article.Title.Eng,
		article.Content.Ru,
		article.Content.Eng,
		pq.Array(keywordsRu),
		pq.Array(keywordsEn),
		article.FilePathId,
		article.VideoPathId,
		pq.Array(article.Literature),
		article.Reference.Ru,
		article.Reference.Eng,
		article.DateReceipt,
		article.DateAcceptance,
		article.DOI,
	)

	var articleID int
	err := row.Scan(&articleID)
	if err != nil {
		return "", err
	}

	for _, author := range article.Authors {
		query := "INSERT INTO authors (fullname_ru, fullname_en, affiliation, email) VALUES ($1, $2, $3, $4) RETURNING author_id"
		row = as.db.QueryRow(
			query,
			author.FullName.Ru,
			author.FullName.Eng,
			author.Affilation,
			author.Email,
		)

		var authorID int
		err := row.Scan(&authorID)
		if err != nil {
			return "", err
		}

		_, err = as.db.Exec(
			"INSERT INTO authors_articles (article, author) VALUES ($1, $2)",
			articleID,
			authorID,
		)
		if err != nil {
			return "", err
		}
	}

	return strconv.Itoa(articleID), nil
}

func (as *ArticleStorage) Find(editionID string) ([]*models.ArticleRead, error) {
	query := fmt.Sprintf(
		"SELECT article_id, %s FROM %s WHERE edition_id = $1",
		getColumns(),
		as.table,
	)
	rows, err := as.db.Query(query, editionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*models.ArticleRead{}

	for rows.Next() {
		var literatureArray pq.StringArray
		var article models.ArticleRead
		var keywordsRu pq.StringArray
		var keywordsEn pq.StringArray

		err := rows.Scan(
			&article.Id,
			&article.EditionId,
			&article.Title.Ru,
			&article.Title.Eng,
			&article.Content.Ru,
			&article.Content.Eng,
			&keywordsRu,
			&keywordsEn,
			&article.FilePathId,
			&article.VideoPathId,
			&literatureArray,
			&article.Reference.Ru,
			&article.Reference.Eng,
			&article.DateReceipt,
			&article.DateAcceptance,
			&article.DOI,
		)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(literatureArray); i++ {
			article.Literature = append(article.Literature, literatureArray[i])
		}

		article.Keywords = make([]models.Text, 0)

		for i := 0; i < len(keywordsRu); i++ {
			keyword := models.Text{
				Ru:  keywordsRu[i],
				Eng: keywordsEn[i],
			}
			article.Keywords = append(article.Keywords, keyword)
		}

		authors, err := as.getAuthorsByArticleID(article.Id)
		if err != nil {
			return nil, err
		}
		article.Authors = authors

		articles = append(articles, &article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (as *ArticleStorage) FindOne(articleIdStr string) (*models.ArticleRead, error) {
	var keywordsRu pq.StringArray
	var keywordsEn pq.StringArray
	var literatureArray pq.StringArray

	query := fmt.Sprintf(
		"SELECT article_id, %s FROM %s WHERE article_id = $1",
		getColumns(),
		as.table,
	)
	row := as.db.QueryRow(query, articleIdStr)

	article := &models.ArticleRead{}
	err := row.Scan(
		&article.Id,
		&article.EditionId,
		&article.Title.Ru,
		&article.Title.Eng,
		&article.Content.Ru,
		&article.Content.Eng,
		&keywordsRu,
		&keywordsEn,
		&article.FilePathId,
		&article.VideoPathId,
		&literatureArray,
		&article.Reference.Ru,
		&article.Reference.Eng,
		&article.DateReceipt,
		&article.DateAcceptance,
		&article.DOI,
	)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(literatureArray); i++ {
		article.Literature = append(article.Literature, literatureArray[i])
	}

	article.Keywords = make([]models.Text, 0)

	for i := 0; i < len(keywordsRu); i++ {
		keyword := models.Text{
			Ru:  keywordsRu[i],
			Eng: keywordsEn[i],
		}
		article.Keywords = append(article.Keywords, keyword)
	}

	authors, err := as.getAuthorsByArticleID(article.Id)
	if err != nil {
		return nil, err
	}
	article.Authors = authors

	return article, nil
}

func (as *ArticleStorage) Delete(IdStr string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE article_id = $1",
		as.table,
	)
	_, err := as.db.Exec(query, IdStr)
	if err != nil {
		return err
	}

	_, err = as.db.Exec("DELETE FROM authors_articles WHERE article = $1", IdStr)
	if err != nil {
		return err
	}
	return err
}

func (as *ArticleStorage) GetFilePathId(idStr string) (string, error) {
	query := fmt.Sprintf(
		"SELECT file_path FROM %s WHERE id = $1",
		as.table,
	)
	row := as.db.QueryRow(query, idStr)

	var filePathId string
	err := row.Scan(&filePathId)
	if err != nil {
		return "", err
	}

	return filePathId, nil
}

func (as *ArticleStorage) GetVideoPathId(idStr string) (string, error) {
	query := fmt.Sprintf(
		"SELECT video_path FROM %s WHERE id = $1",
		as.table,
	)
	row := as.db.QueryRow(query, idStr)

	var filePathId string
	err := row.Scan(&filePathId)
	if err != nil {
		return "", err
	}

	return filePathId, nil
}

func (as *ArticleStorage) UpdateOne(newArticle *models.ArticleUpdate) error {
	var keywordsRu []string
	var keywordsEn []string

	for _, keyword := range newArticle.Keywords {
		keywordsRu = append(keywordsRu, keyword.Ru)
		keywordsEn = append(keywordsEn, keyword.Eng)
	}

	query := `
	UPDATE articles
	SET
	title_ru = $2, 
	title_en = $3, 
	content_ru = $4, 
	content_en = $5, 
	keywords_ru = $6, 
	keywords_en = $7, 
	file_path = $8, 
	video_path = $9, 
	literature = $10, 
	reference_ru = $11, 
	reference_en = $12, 
	date_receipt = $13, 
	date_acceptance = $14, 
	doi = $15
	WHERE
		article_id = $1
`
	_, err := as.db.Exec(
		query,
		&newArticle.Id,
		&newArticle.Title.Ru,
		&newArticle.Title.Eng,
		&newArticle.Content.Ru,
		&newArticle.Content.Eng,
		&keywordsRu,
		&keywordsEn,
		&newArticle.FilePathId,
		&newArticle.VideoPathId,
		&newArticle.Literature,
		&newArticle.Reference.Ru,
		&newArticle.Reference.Eng,
		&newArticle.DateReceipt,
		&newArticle.DateAcceptance,
		&newArticle.DOI,
	)
	if err != nil {
		return err
	}

	_, err = as.db.Exec("DELETE FROM authors_articles WHERE article = $1", newArticle.Id)
	if err != nil {
		return err
	}

	for _, author := range newArticle.Authors {
		row := as.db.QueryRow("INSERT INTO authors (fullname_ru, fullname_en, affiliation, email) VALUES ($1, $2, $3, $4) RETURNING author_id",
			author.FullName.Ru,
			author.FullName.Eng,
			author.Affilation,
			author.Email,
		)

		var authorID int
		err = row.Scan(&authorID)
		if err != nil {
			return err
		}

		_, err = as.db.Exec("INSERT INTO authors_articles (article, author) VALUES ($1, $2)", newArticle.Id, authorID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (as *ArticleStorage) getAuthorsByArticleID(articleID int) ([]models.Author, error) {
	query := "SELECT fullname_ru, fullname_en, affiliation, email FROM authors WHERE author_id IN (SELECT author FROM authors_articles WHERE article = $1)"
	rows, err := as.db.Query(query, articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []models.Author

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
