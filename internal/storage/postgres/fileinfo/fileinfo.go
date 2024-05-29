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
		"INSERT INTO %s (filename, bucket) VALUES ($1, $2) RETURNING file_id",
		f.table,
	)

	row := f.db.QueryRow(
		query,
		fileInfo.Filename,
		fileInfo.BucketName,
	)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (f *FileInfoStorage) DeleteOne(id string) error {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE edition_id = $1",
		f.table,
	)

	_, err := f.db.Exec(
		query,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (f *FileInfoStorage) FindOne(id string) (*models.FileInfo, error) {
	query := fmt.Sprintf(
		"SELECT filename, bucket FROM %s WHERE file_id = $1",
		f.table,
	)

	row := f.db.QueryRow(query, id)

	fileInfo := &models.FileInfo{}
	err := row.Scan(&fileInfo.Filename, &fileInfo.BucketName)
	if err != nil {
		return nil, err
	}
	return fileInfo, nil
}
