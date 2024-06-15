package commentst

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/internal/storage/postgres/utils"
)

type CommentsStorage struct {
	db    *sql.DB
	table string
}

func NewCommentStorage(db *sql.DB, table string) *CommentsStorage {
	return &CommentsStorage{
		db:    db,
		table: table,
	}
}

func getColumns() string {
	return "article, content_ru, content_en, date_create, is_approved, author"
}

func generateValuesID(count int) string {
	res := ""
	for i := 1; i < count; i++ {
		res += fmt.Sprintf("$%v, ", i)
	}
	res += fmt.Sprintf("$%v", count)
	return res
}

func (cs *CommentsStorage) InsertOne(comment *models.CommentCreate) (string, error) {
	var id string
	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s) RETURNING comment_id",
		cs.table,
		getColumns(),
		generateValuesID(6),
	)

	row := cs.db.QueryRow(
		query,
		comment.ArticleId,
		comment.Content.Ru,
		comment.Content.Eng,
		comment.Date,
		comment.IsApproved,
		comment.Author)

	err := row.Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (cs *CommentsStorage) GetAll(args *models.CommentQuery) ([]*models.CommentRead, error) {
	var comments []*models.CommentRead
	query := fmt.Sprintf(
		"SELECT comment_id, %s FROM %s WHERE article = $1 AND is_approved = $2",
		getColumns(),
		cs.table,
	)

	queryBatch := utils.AddBatchToQuery(query, &args.BatchArgs)

	rows, err := cs.db.Query(
		queryBatch,
		args.ArticleID,
		args.OnlyApproved,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.CommentRead
		err := rows.Scan(
			&comment.Id,
			&comment.ArticleId,
			&comment.Content.Ru,
			&comment.Content.Eng,
			&comment.Date,
			&comment.IsApproved,
			&comment.Author)

		if err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}

func (cs *CommentsStorage) UpdateOne(comment *models.CommentUpdate) error {
	query := fmt.Sprintf(
		"UPDATE %s SET content_ru = $1, content_en = $2 WHERE comment_id = $3",
		cs.table,
	)

	_, err := cs.db.Exec(
		query,
		comment.Content.Ru,
		comment.Content.Eng,
		comment.Id)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CommentsStorage) Delete(idStr string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE comment_id = $1",
		cs.table,
	)

	_, err := cs.db.Exec(query, idStr)

	if err != nil {
		return err
	}

	return nil
}

func (cs *CommentsStorage) Approve(commentApprove *models.CommentApprove, contentField string) error {
	query := fmt.Sprintf(
		"UPDATE %s SET is_approved = true, "+contentField+" = $1 WHERE comment_id = $2 AND is_approved = false",
		cs.table,
	)

	_, err := cs.db.Exec(
		query,
		commentApprove.ContentPart,
		commentApprove.Id,
	)

	if err != nil {
		return err
	}

	return nil
}
