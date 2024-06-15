package articlest

import (
	"database/sql"
	"fmt"
	"strconv"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/internal/storage/postgres/utils"

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
		article.DocumentID,
		article.VideoID,
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

	return strconv.Itoa(articleID), nil
}

func (as *ArticleStorage) Find(args *models.ArticleQuery) ([]*models.ArticleRead, error) {
	query := fmt.Sprintf(
		"SELECT article_id, %s FROM %s WHERE edition_id = $1",
		getColumns(),
		as.table,
	)

	queryBatch := utils.AddBatchToQuery(query, &args.BatchArgs)

	fmt.Println(query)

	rows, err := as.db.Query(queryBatch, args.EditionID)
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
			&article.DocumentID,
			&article.VideoID,
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
		&article.DocumentID,
		&article.VideoID,
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

	return err
}

func (as *ArticleStorage) GetDocumentID(idStr string) (string, error) {
	query := fmt.Sprintf(
		"SELECT file_path FROM %s WHERE id = $1",
		as.table,
	)
	row := as.db.QueryRow(query, idStr)

	var documentID string
	err := row.Scan(&documentID)
	if err != nil {
		return "", err
	}

	return documentID, nil
}

func (as *ArticleStorage) GetVideoID(idStr string) (string, error) {
	query := fmt.Sprintf(
		"SELECT video_path FROM %s WHERE id = $1",
		as.table,
	)
	row := as.db.QueryRow(query, idStr)

	var videoID string
	err := row.Scan(&videoID)
	if err != nil {
		return "", err
	}

	return videoID, nil
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
		pq.Array(&keywordsRu),
		pq.Array(&keywordsEn),
		&newArticle.DocumentID,
		&newArticle.VideoID,
		pq.Array(&newArticle.Literature),
		&newArticle.Reference.Ru,
		&newArticle.Reference.Eng,
		&newArticle.DateReceipt,
		&newArticle.DateAcceptance,
		&newArticle.DOI,
	)
	if err != nil {
		return err
	}

	return nil
}
