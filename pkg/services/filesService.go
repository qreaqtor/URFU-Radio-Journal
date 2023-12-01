package services

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"urfu-radio-journal/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type FilesService struct {
	ctx         *context.Context
	storage     *mongo.Collection
	basePath    string
	directories map[string]string
}

func NewFilesService() *FilesService {
	return &FilesService{
		ctx:         db.GetContext(),
		storage:     db.GetStorage("files"),
		basePath:    "../attachments",
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

func (this *FilesService) GetFilePath(filename, resourceType string) (path string, err error) {
	ext := filepath.Ext(filename)
	if dir, ok := this.directories[ext]; ok {
		path = fmt.Sprintf("%s/%s/%s/%s%s", this.basePath, resourceType, dir, filename, ext)
	} else {
		err = errors.New("This file extension is not supported.")
	}
	return
}

func (this *FilesService) CheckFilePath(filename, resourceType string) (path string, err error) {
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

func (this *FilesService) GetFileURL(filename, resourceType, path, identifierStr string) (url string, err error) {
	var res *mongo.InsertOneResult
	var identifier primitive.ObjectID
	if identifierStr != "" {
		err = os.Remove(path)
		if err != nil {
			return
		}
		identifier, err = primitive.ObjectIDFromHex(identifierStr)
		if err != nil {
			return
		}
		filter := bson.M{"_id": identifier}
		update := bson.M{"$set": bson.M{"filename": filename}}
		_, err = this.storage.UpdateOne(*this.ctx, filter, update)
	} else {
		res, err = this.storage.InsertOne(*this.ctx, bson.M{"filename": filename})
	}
	ext := filepath.Ext(filename)
	url = fmt.Sprintf("/files/download/%s/%s%s", resourceType, res.InsertedID, ext)
	return
}
