package buckets

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
}

type AllowedContentType map[FileRepo][]string

type Buckets struct {
	bucketsByContentType, bucketsbyName map[string]FileRepo
}

func NewBuckets(types AllowedContentType) *Buckets {
	bucketsByContentType := make(map[string]FileRepo)
	bucketsbyName := make(map[string]FileRepo)

	for bucket, contentTypes := range types {
		bucketsbyName[bucket.GetBucketName()] = bucket
		for _, contentType := range contentTypes {
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
