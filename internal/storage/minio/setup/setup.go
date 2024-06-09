package miniost

import (
	"context"
	"fmt"
	"urfu-radio-journal/internal/config"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func GetConnect(conf config.MinioConfig, ssl bool) (*minio.Client, error) {
	client, err := minio.New(conf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(conf.User, conf.Password, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func InitBuckets(ctx context.Context, client *minio.Client, buckets ...config.BucketConfig) error {
	opts := minio.MakeBucketOptions{}
	for _, bucket := range buckets {
		exists, err := client.BucketExists(ctx, bucket.Name)
		if err != nil {
			return fmt.Errorf("cant check is bucket exists: %v", err)
		}
		if !exists {
			err = client.MakeBucket(ctx, bucket.Name, opts)
			if err != nil {
				return fmt.Errorf("cant make bucket: %v", err)
			}
		}
	}
	return nil
}
