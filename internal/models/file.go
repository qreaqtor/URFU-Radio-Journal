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
	BacketName string // this field sets in service
}
