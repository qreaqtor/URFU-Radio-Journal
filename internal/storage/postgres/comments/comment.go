package commentst

import (
	"database/sql"
	"urfu-radio-journal/internal/models"
)

type CommentsStorage struct {
	db *sql.DB
}

func NewCommentStorage(db *sql.DB) *CommentsStorage {
	return &CommentsStorage{
		db: db,
	}
}

func (cs *CommentsStorage) InsertOne(comment *models.CommentCreate) (string, error) {
	stmt, err := cs.db.Prepare("INSERT INTO comments (article, content_ru, content_en, date_create, is_approved, author) VALUES ($1, $2, $3, $4, $5, $6) RETURNING comment_id")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var id string
	err = stmt.QueryRow(comment.ArticleId, comment.Content.Ru, comment.Content.Eng, comment.Date, comment.IsApproved, comment.Author).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (cs *CommentsStorage) GetAll(onlyApproved bool, articleIdStr string) ([]*models.CommentRead, error) {
	rows, err := cs.db.Query(
		"SELECT comment_id, article, content_ru, content_en, date_create, is_approved, author FROM comments WHERE article = $1 AND is_approved = $2", 
		articleIdStr, 
		onlyApproved,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []*models.CommentRead
	for rows.Next() {
		var comment models.CommentRead
		err := rows.Scan(&comment.Id, &comment.ArticleId, &comment.Content.Ru, &comment.Content.Eng, &comment.Date, &comment.IsApproved, &comment.Author)
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
	_, err := cs.db.Exec("UPDATE comments SET content_ru = $1, content_en = $2 WHERE comment_id = $3", comment.Content.Ru, comment.Content.Eng, comment.Id)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CommentsStorage) Delete(idStr string) error {
	_, err := cs.db.Exec("DELETE FROM comments WHERE comment_id = $1", idStr)
	if err != nil {
		return err
	}
	return nil
}

func (cs *CommentsStorage) Approve(commentApprove *models.CommentApprove, contentField string) error {
	_, err := cs.db.Exec(
		"UPDATE comments SET is_approved = true, "+contentField+" = $1 WHERE comment_id = $2 AND is_approved = false", 
		commentApprove.ContentPart, 
		commentApprove.Id,
	)
	if err != nil {
		return err
	}
	return nil
}
