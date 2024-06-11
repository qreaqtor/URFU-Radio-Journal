package filesrv

import (
	"context"
	"fmt"
	"urfu-radio-journal/internal/models"
	bucketsrv "urfu-radio-journal/internal/services/files/buckets"
)

type fileInfoRepo interface {
	InsertOne(string) (string, error)
	DeleteOne(string) (string, error)
	FindOne(string) (string, error)
}

type FileService struct {
	filesInfo fileInfoRepo
	buckets   *bucketsrv.Buckets
}

func NewFileService(files fileInfoRepo, buckets ...bucketsrv.FileRepo) *FileService {
	return &FileService{
		filesInfo: files,
		buckets:   bucketsrv.NewBuckets(buckets),
	}
}

func (f *FileService) UploadFile(ctx context.Context, fileUnit *models.FileUnit) (string, error) {
	bucket, err := f.buckets.GetBucketByContentType(fileUnit.ContentType)
	if err != nil {
		return "", err
	}

	id, err := f.filesInfo.InsertOne(bucket.GetBucketName())
	if err != nil {
		return "", err
	}

	err = bucket.UploadFile(ctx, fileUnit, id)
	if err != nil {
		_, err2 := f.filesInfo.DeleteOne(id)
		if err2 != nil {
			return "", fmt.Errorf("%v; %v", err, err2)
		}
		return "", err
	}

	return id, nil
}

func (f *FileService) DownloadFile(ctx context.Context, id string) (*models.FileUnit, error) {
	bucketName, err := f.filesInfo.FindOne(id)
	if err != nil {
		return nil, err
	}

	bucket, err := f.buckets.GetBucketByName(bucketName)
	if err != nil {
		return nil, err
	}

	fileUnit, err := bucket.DownloadFile(ctx, id)
	if err != nil {
		return nil, err
	}

	return fileUnit, nil
}

func (f *FileService) DeleteFile(ctx context.Context, id string) error {
	bucketName, err := f.filesInfo.DeleteOne(id)
	if err != nil {
		return err
	}

	bucket, err := f.buckets.GetBucketByName(bucketName)
	if err != nil {
		return err
	}

	return bucket.DeleteFile(ctx, id)
}
