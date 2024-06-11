package fileinfost

import (
	"database/sql"
	"fmt"
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

func (f *FileInfoStorage) InsertOne(bucketName string) (string, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (bucket) VALUES ($1) RETURNING file_id",
		f.table,
	)

	row := f.db.QueryRow(query, bucketName)

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (f *FileInfoStorage) DeleteOne(id string) (string, error) {
	query := fmt.Sprintf(
		"DELETE FROM %s WHERE file_id = $1 RETURNING bucket",
		f.table,
	)

	row := f.db.QueryRow(query, id)

	var bucket string
	err := row.Scan(&bucket)
	if err != nil {
		return "", err
	}

	return bucket, nil
}

func (f *FileInfoStorage) FindOne(id string) (string, error) {
	query := fmt.Sprintf(
		"SELECT bucket FROM %s WHERE file_id = $1",
		f.table,
	)

	row := f.db.QueryRow(query, id)

	var bucketName string
	err := row.Scan(&bucketName)
	if err != nil {
		return "", err
	}
	return bucketName, nil
}
