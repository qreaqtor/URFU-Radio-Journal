package models

import "io"

type FileUnit struct {
	Info      *FileInfo
	Payload   io.Reader
	PayloadID string
}

type FileInfo struct {
	Name        string
	ContentType string
	Size        int64
}
