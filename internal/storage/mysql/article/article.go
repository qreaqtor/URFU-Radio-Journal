package articlest

import (
	"database/sql"
	"strconv"
	"urfu-radio-journal/internal/models"

	"github.com/lib/pq"
)

type ArticleStorage struct {
	db *sql.DB
}

func NewArticleStorage(db *sql.DB) *ArticleStorage {
	return &ArticleStorage{
		db: db,
	}
}

func (a *ArticleStorage) InsertOne(article *models.ArticleCreate) (string, error) {
	var keywordsRu []string
	var keywordsEn []string

	for _, keyword := range article.Keywords {
		keywordsRu = append(keywordsRu, keyword.Ru)
		keywordsEn = append(keywordsEn, keyword.Eng)
	}

	var articleID int
	err := a.db.QueryRow("INSERT INTO articles (edition_id, title_ru, title_en, content_ru, content_en, keywords_ru, keywords_en, file_path, video_path, literature, reference_ru, reference_en, date_receipt, date_acceptance, doi) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING article_id;",
		article.EditionId, article.Title.Ru, article.Title.Eng, article.Content.Ru, article.Content.Eng,
		pq.Array(keywordsRu), pq.Array(keywordsEn), article.FilePathId, article.VideoPathId, pq.Array(article.Literature),
		article.Reference.Ru, article.Reference.Eng, article.DateReceipt, article.DateAcceptance, article.DOI).Scan(&articleID)

	if err != nil {
		return "", err
	}

	for _, author := range article.Authors {
		var authorID int
		err = a.db.QueryRow("INSERT INTO authors (fullname_ru, fullname_en, affiliation, email) VALUES ($1, $2, $3, $4) RETURNING author_id",
			author.FullName.Ru, author.FullName.Eng, author.Affilation, author.Email).Scan(&authorID)
		if err != nil {
			return "", err
		}

		_, err = a.db.Exec(
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
	rows, err := as.db.Query("SELECT article_id, edition_id, title_ru, title_en, content_ru, content_en, keywords_ru, keywords_en, file_path, video_path, literature, reference_ru, reference_en, date_receipt, date_acceptance, doi FROM articles WHERE edition_id = $1", editionID)
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
	article := &models.ArticleRead{}
	var keywordsRu pq.StringArray
	var keywordsEn pq.StringArray
	var literatureArray pq.StringArray

	err := as.db.QueryRow("SELECT article_id, edition_id, title_ru, title_en, content_ru, content_en, keywords_ru, keywords_en, file_path, video_path, literature, reference_ru, reference_en, date_receipt, date_acceptance, doi FROM articles WHERE article_id = $1", articleIdStr).
		Scan(
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
	_, err := as.db.Exec("DELETE FROM articles WHERE article_id = $1", IdStr)
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
	var filePathId string

	err := as.db.QueryRow("SELECT file_path FROM articles WHERE id = $1", idStr).Scan(&filePathId)
	if err != nil {
		return "", err
	}

	return filePathId, nil
}

func (as *ArticleStorage) GetVideoPathId(idStr string) (string, error) {
	var filePathId string

	err := as.db.QueryRow("SELECT video_path FROM articles WHERE id = $1", idStr).Scan(&filePathId)
	if err != nil {
		return "", err
	}

	return filePathId, nil
}

func (as *ArticleStorage) UpdateOne(newArticle *models.ArticleUpdate) error {
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

	var keywordsRu []string
	var keywordsEn []string

	for _, keyword := range newArticle.Keywords {
		keywordsRu = append(keywordsRu, keyword.Ru)
		keywordsEn = append(keywordsEn, keyword.Eng)
	}

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
		var authorID int
		err = as.db.QueryRow("INSERT INTO authors (fullname_ru, fullname_en, affiliation, email) VALUES ($1, $2, $3, $4) RETURNING author_id",
			author.FullName.Ru, author.FullName.Eng, author.Affilation, author.Email).Scan(&authorID)
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
	rows, err := as.db.Query("SELECT fullname_ru, fullname_en, affiliation, email FROM authors WHERE author_id IN (SELECT author FROM authors_articles WHERE article = $1)", articleID)
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
