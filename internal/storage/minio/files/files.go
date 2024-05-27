package filest

import (
	"context"
	"fmt"
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
	info, err := f.client.PutObject(
		ctx,
		f.bucket,
		file.PayloadID,
		file.Payload,
		file.Info.Size,
		minio.PutObjectOptions{},
	)
	if err != nil {
		return err
	}
	fmt.Println(info.Key)
	return nil
}

func (f *FileStorage) DeleteFile(ctx context.Context, objName string) error {
	err := f.client.RemoveObject(
		ctx,
		f.bucket,
		objName,
		minio.RemoveObjectOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}
