package models

import "io"

type FileUnit struct {
	Payload     io.ReadCloser
	Name        string // this field sets in service
	Size        int64
	ContentType string
}
