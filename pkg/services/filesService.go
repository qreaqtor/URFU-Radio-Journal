package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type FilesService struct {
	basePath    string
	directories map[string]string
}

func NewFilesService() *FilesService {
	path := "../attachments"
	return &FilesService{
		basePath:    path,
		directories: getDirs(),
	}
}

func getDirs() map[string]string {
	dirs := make(map[string]string, 3)
	dirs[".pdf"] = "documents"
	dirs[".png"] = "images"
	dirs[".mkv"] = "videos"
	return dirs
}

func (this *FilesService) GetFilePath(filename string, resourceId string, resourceType string) (path string, err error) {
	ext := filepath.Ext(filename)
	if dir, ok := this.directories[ext]; ok {
		path = fmt.Sprintf("%s/%s/%s/%s%s", this.basePath, resourceType, dir, resourceId, ext)
	} else {
		err = errors.New("This file extension is not supported.")
	}
	return
}

func (this *FilesService) CheckFilePath(filename string, resourceType string) (path string, err error) {
	ext := filepath.Ext(filename)
	if dir, ok := this.directories[ext]; ok {
		path = fmt.Sprintf("%s/%s/%s/%s", this.basePath, resourceType, dir, filename)
		_, err := os.Stat(path)
		if err != nil && os.IsNotExist(err) {
			err = errors.New("This file not exist.")
		}
	} else {
		err = errors.New("This file extension is not supported.")
	}
	return
}
