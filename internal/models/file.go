package models

import "io"

type FileUnit struct {
	Payload     io.Reader
	PayloadName string
	PayloadSize int64
}
