package models

import "io"

type FileUnit struct {
	Payload     io.Reader
	InfoID      string // this field sets in service
	Size        int64
	ContentType string
}

type FileInfo struct {
	Filename   string
	BucketName string // this field sets in service
}
