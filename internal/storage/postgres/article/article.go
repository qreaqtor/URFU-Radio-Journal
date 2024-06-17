package articlest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/internal/utils"

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
	return "edition_id, title, content, keywords, file_path, video_path, literature, reference, date_receipt, date_acceptance, doi, authors"
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
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING article_id;",
		as.table,
		getColumns(),
		generateValuesID(12),
	)

	jsonTitle, err := json.Marshal(article.Title)
	if err != nil {
		return "", err
	}

	jsonContent, err := json.Marshal(article.Content)
	if err != nil {
		return "", err
	}

	jsonKeywords, err := json.Marshal(article.Keywords)
	if err != nil {
		return "", err
	}

	jsonReference, err := json.Marshal(article.Reference)
	if err != nil {
		return "", err
	}

	jsonAuthors, err := json.Marshal(article.Authors)
	if err != nil {
		return "", err
	}

	row := as.db.QueryRow(
		query,
		article.EditionId,
		jsonTitle,
		jsonContent,
		jsonKeywords,
		article.DocumentID,
		article.VideoID,
		pq.Array(article.Literature),
		jsonReference,
		article.DateReceipt,
		article.DateAcceptance,
		article.DOI,
		jsonAuthors,
	)

	var articleID int
	err = row.Scan(&articleID)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(articleID), nil
}

func (as *ArticleStorage) Find(args *models.ArticleQuery) ([]*models.ArticleRead, error) {
	var title, content, reference models.Text
	keywords := []models.Text{}
	authors := []models.Author{}

	var jsonTitle, jsonContent, jsonKeywords, jsonReference, jsonAuthors []byte

	query := fmt.Sprintf(
		"SELECT article_id, %s FROM %s",
		getColumns(),
		as.table,
	)

	query, isSearch := utils.AddSearchToQuery(query, args.ArticleSearch)

	query = utils.AddBatchToQuery(query, &args.BatchArgs)

	search := fmt.Sprint(args.EditionID)
	if isSearch {
		search = strings.ReplaceAll(args.Search, " ", " | ")
	}

	fmt.Println(query)

	rows, err := as.db.Query(query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []*models.ArticleRead{}

	for rows.Next() {
		var literatureArray pq.StringArray
		var article models.ArticleRead

		err := rows.Scan(
			&article.Id,
			&article.EditionId,
			&jsonTitle,
			&jsonContent,
			&jsonKeywords,
			&article.DocumentID,
			&article.VideoID,
			&literatureArray,
			&jsonReference,
			&article.DateReceipt,
			&article.DateAcceptance,
			&article.DOI,
			&jsonAuthors,
		)
		if err != nil {
			return nil, err
		}

		for i := 0; i < len(literatureArray); i++ {
			article.Literature = append(article.Literature, literatureArray[i])
		}

		err = json.Unmarshal(jsonTitle, &title)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(jsonContent, &content)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(jsonKeywords, &keywords)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(jsonReference, &reference)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(jsonAuthors, &authors)
		if err != nil {
			return nil, err
		}

		article.Title = title
		article.Content = content
		article.Keywords = keywords
		article.Reference = reference
		article.Authors = authors

		articles = append(articles, &article)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return articles, nil
}

func (as *ArticleStorage) FindOne(articleIdStr string) (*models.ArticleRead, error) {
	var title, content, reference models.Text
	keywords := []models.Text{}
	authors := []models.Author{}

	var jsonTitle, jsonContent, jsonKeywords, jsonReference, jsonAuthors []byte

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
		&jsonTitle,
		&jsonContent,
		&jsonKeywords,
		&article.DocumentID,
		&article.VideoID,
		&literatureArray,
		&jsonReference,
		&article.DateReceipt,
		&article.DateAcceptance,
		&article.DOI,
		&jsonAuthors,
	)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(literatureArray); i++ {
		article.Literature = append(article.Literature, literatureArray[i])
	}

	err = json.Unmarshal(jsonTitle, &title)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonContent, &content)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonKeywords, &keywords)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonReference, &reference)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonAuthors, &authors)
	if err != nil {
		return nil, err
	}

	article.Title = title
	article.Content = content
	article.Keywords = keywords
	article.Reference = reference
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
	query := `
	UPDATE articles
	SET
	title = $2, 
	content = $3, 
	keywords = $4, 
	file_path = $5, 
	video_path = $6, 
	literature = $7, 
	reference = $8, 
	date_receipt = $9, 
	date_acceptance = $10, 
	doi = $11,
	authors = $12
	WHERE
		article_id = $1
`

	jsonTitle, err := json.Marshal(newArticle.Title)
	if err != nil {
		return err
	}

	jsonContent, err := json.Marshal(newArticle.Content)
	if err != nil {
		return err
	}

	jsonKeywords, err := json.Marshal(newArticle.Keywords)
	if err != nil {
		return err
	}

	jsonReference, err := json.Marshal(newArticle.Reference)
	if err != nil {
		return err
	}

	jsonAuthors, err := json.Marshal(newArticle.Authors)
	if err != nil {
		return err
	}

	_, err = as.db.Exec(
		query,
		&newArticle.Id,
		&jsonTitle,
		&jsonContent,
		&jsonKeywords,
		&newArticle.DocumentID,
		&newArticle.VideoID,
		pq.Array(&newArticle.Literature),
		&jsonReference,
		&newArticle.DateReceipt,
		&newArticle.DateAcceptance,
		&newArticle.DOI,
		&jsonAuthors,
	)
	if err != nil {
		return err
	}

	return nil
}

func (as *ArticleStorage) GetCount() (int, error) {
	query := fmt.Sprintf(
		"SELECT COUNT(*) FROM %s",
		as.table,
	)

	row := as.db.QueryRow(query)

	count := 0
	err := row.Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
