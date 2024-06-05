package models

import "io"

type FileUnit struct {
	Payload     io.ReadCloser
	Size        int64
	ContentType string
}
