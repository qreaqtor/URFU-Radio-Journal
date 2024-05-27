package fileinfost

import (
	"database/sql"
	"fmt"
	"urfu-radio-journal/internal/models"
)

type FileInfoStorage struct {
	db    *sql.DB
	table string
}

func NewFileInfoStorage(db *sql.DB, table string) *FileInfoStorage {
	return &FileInfoStorage{
		db:    db,
		table: table,
	}
}

func (f *FileInfoStorage) InsertOne(fileInfo *models.FileInfo) (string, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (filename, content_type, size) VALUES ($1, $2, $3) RETURNING file_id",
		f.table,
	)

	row := f.db.QueryRow(
		query,
		fileInfo.Name,
		fileInfo.ContentType,
		fileInfo.Size,
	)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (f *FileInfoStorage) DeleteOne(id string) error {
	return nil
}

func (f *FileInfoStorage) FindOne(id string) (*models.FileInfo, error) {
	return nil, nil
}

func (f *FileInfoStorage) UpdateOne(id string) error {
	return nil
}
