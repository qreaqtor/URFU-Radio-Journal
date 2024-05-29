package filest

import (
	"context"
	"urfu-radio-journal/internal/models"

	"github.com/minio/minio-go/v7"
)

type FileStorage struct {
	client *minio.Client
	bucket string
}

func NewFileStorage(client *minio.Client, bucketName string) *FileStorage {
	return &FileStorage{
		client: client,
		bucket: bucketName,
	}
}

func (f *FileStorage) UploadFile(ctx context.Context, file *models.FileUnit) error {
	_, err := f.client.PutObject(
		ctx,
		f.bucket,
		file.Name,
		file.Payload,
		file.Size,
		minio.PutObjectOptions{ContentType: file.ContentType},
	)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStorage) DeleteFile(ctx context.Context, id string) error {
	err := f.client.RemoveObject(
		ctx,
		f.bucket,
		id,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileStorage) DownloadFile(ctx context.Context, id string) (*models.FileUnit, error) {
	obj, err := f.client.GetObject(
		ctx,
		f.bucket,
		id,
		minio.GetObjectOptions{},
	)
	if err != nil {
		return nil, err
	}

	stat, err := obj.Stat()
	if err != nil {
		return nil, err
	}

	file := &models.FileUnit{
		Name:        id,
		Payload:     obj,
		Size:        stat.Size,
		ContentType: stat.ContentType,
	}

	return file, nil
}

func (f *FileStorage) GetBucketName() string {
	return f.bucket
}
