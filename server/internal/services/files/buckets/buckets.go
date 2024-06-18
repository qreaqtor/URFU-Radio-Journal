package bucketsrv

import (
	"context"
	"errors"
	"urfu-radio-journal/internal/models"
)

var (
	errBadContentType = errors.New("bad Content-Type")
	errBadName        = errors.New("unknown bucket name")
)

type FileRepo interface {
	UploadFile(context.Context, *models.FileUnit, string) error
	DeleteFile(context.Context, string) error
	DownloadFile(context.Context, string) (*models.FileUnit, error)
	GetBucketName() string
	GetContentTypes() []string
}

type Buckets struct {
	bucketsByContentType, bucketsbyName map[string]FileRepo
}

func NewBuckets(buckets []FileRepo) *Buckets {
	bucketsByContentType := make(map[string]FileRepo)
	bucketsbyName := make(map[string]FileRepo)

	for _, bucket := range buckets {
		bucketsbyName[bucket.GetBucketName()] = bucket
		for _, contentType := range bucket.GetContentTypes() {
			bucketsByContentType[contentType] = bucket
		}
	}

	return &Buckets{
		bucketsByContentType: bucketsByContentType,
		bucketsbyName:        bucketsbyName,
	}
}

func (b *Buckets) GetBucketByContentType(contentType string) (FileRepo, error) {
	if bucket, ok := b.bucketsByContentType[contentType]; ok {
		return bucket, nil
	}
	return nil, errBadContentType
}

func (b *Buckets) GetBucketByName(name string) (FileRepo, error) {
	if bucket, ok := b.bucketsbyName[name]; ok {
		return bucket, nil
	}
	return nil, errBadName
}
