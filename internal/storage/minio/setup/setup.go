package miniost

import (
	"context"
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func GetConnect(user, password, url string, ssl bool) (*minio.Client, error) {
	client, err := minio.New(url, &minio.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: ssl,
	})
	if err != nil {
		return nil, err
	}

	return client, nil
}

func InitBuckets(ctx context.Context, client *minio.Client, buckets ...string) error {
	opts := minio.MakeBucketOptions{}
	for _, bucket := range buckets {
		exists, err := client.BucketExists(ctx, bucket)
		if err != nil {
			return fmt.Errorf("cant check bucket exists: %v", err)
		}
		if !exists {
			err = client.MakeBucket(ctx, bucket, opts)
			if err != nil {
				return fmt.Errorf("cant make bucket: %v", err)
			}
		}
	}
	return nil
}
