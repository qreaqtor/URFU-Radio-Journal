package filest

import (
	"context"
	"urfu-radio-journal/internal/models"

	"github.com/minio/minio-go/v7"
)

const (
	filenameKey = "Filename"
)

type FileStorage struct {
	client       *minio.Client
	bucket       string
	contentTypes []string
}

func NewFileStorage(client *minio.Client, bucketName string, content ...string) *FileStorage {
	return &FileStorage{
		client:       client,
		bucket:       bucketName,
		contentTypes: content,
	}
}

func (f *FileStorage) UploadFile(ctx context.Context, file *models.FileUnit, id string) error {
	_, err := f.client.PutObject(
		ctx,
		f.bucket,
		id,
		file.Payload,
		file.Size,
		minio.PutObjectOptions{
			ContentType: file.ContentType,
			UserMetadata: map[string]string{
				filenameKey: file.Filename,
			},
		},
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
		Payload:     obj,
		Size:        stat.Size,
		ContentType: stat.ContentType,
		Filename:    stat.UserMetadata[filenameKey],
	}

	return file, nil
}

func (f *FileStorage) GetBucketName() string {
	return f.bucket
}

func (f *FileStorage) GetContentTypes() []string {
	return f.contentTypes
}
