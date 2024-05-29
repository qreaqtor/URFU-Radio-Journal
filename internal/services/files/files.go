package filesrv

import (
	"context"
	"fmt"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/internal/services/files/buckets"
)

type fileInfoRepo interface {
	InsertOne(*models.FileInfo) (string, error)
	DeleteOne(string) error
	FindOne(string) (*models.FileInfo, error)
}

type FileService struct {
	filesInfo fileInfoRepo
	buckets   *buckets.Buckets
}

func NewFileService(types buckets.AllowedContentType, files fileInfoRepo) *FileService {
	buckets := buckets.NewBuckets(types)
	return &FileService{
		filesInfo: files,
		buckets:   buckets,
	}
}

func (f *FileService) UploadFile(ctx context.Context, fileUnit *models.FileUnit, fileInfo *models.FileInfo) (string, error) {
	bucket, err := f.buckets.GetBucketByContentType(fileUnit.ContentType)
	if err != nil {
		return "", err
	}
	fileInfo.BucketName = bucket.GetBucketName()

	id, err := f.filesInfo.InsertOne(fileInfo)
	if err != nil {
		return "", err
	}

	fileUnit.Name = id

	err = bucket.UploadFile(ctx, fileUnit)
	if err != nil {
		err2 := f.filesInfo.DeleteOne(id)
		if err2 != nil {
			return "", fmt.Errorf("%v; %v", err, err2)
		}
		return "", err
	}

	return fileUnit.Name, nil
}

func (f *FileService) DownloadFile(ctx context.Context, id string) (*models.FileUnit, error) {
	info, err := f.filesInfo.FindOne(id)
	if err != nil {
		return nil, err
	}

	bucket, err := f.buckets.GetBucketByName(info.BucketName)
	if err != nil {
		return nil, err
	}

	fileUnit, err := bucket.DownloadFile(ctx, id)
	if err != nil {
		return nil, err
	}

	fileUnit.Name = info.Filename

	return fileUnit, nil
}

func (f *FileService) DeleteFile(ctx context.Context, id string) error {
	info, err := f.filesInfo.FindOne(id)
	if err != nil {
		return err
	}

	bucket, err := f.buckets.GetBucketByName(info.BucketName)
	if err != nil {
		return err
	}

	return bucket.DeleteFile(ctx, id)
}

// func (f *FileService) getBucketByID(id string) (fileRepo, error) {

// 	bucket, err := f.getBucketByContentType(info.ContentType)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return bucket, err
// }
