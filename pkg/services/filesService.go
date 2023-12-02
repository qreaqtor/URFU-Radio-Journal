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
		storage:     db.GetStorage("filePaths"),
		basePath:    "../attachments",
		directories: getDirs(),
	}
}

func getDirs() map[string]string {
	dirs := make(map[string]string, 3)
	dirs[".pdf"] = "documents"
	dirs[".png"] = "images"
	dirs[".mkv"] = "videos"
	dirs[".mp4"] = "videos"
	return dirs
}

func (this *FilesService) CheckFilePath(filePathIdStr, resourceType string) (path string, err error) {
	filePathId, err := primitive.ObjectIDFromHex(filePathIdStr)
	if err != nil {
		return
	}
	var filePath struct {
		Id   primitive.ObjectID
		Path string
	}
	filter := bson.M{"_id": filePathId}
	err = this.storage.FindOne(*this.ctx, filter).Decode(&filePath)
	path = filePath.Path
	return
}

func (this *FilesService) GetFileURL(filename, resourceType, filePathIdStr string) (filePathId primitive.ObjectID, path string, err error) {
	if path, err = this.getFilePath(filename, resourceType); err != nil {
		return
	}
	if filePathIdStr != "" {
		filePathId, err = primitive.ObjectIDFromHex(filePathIdStr)
		if err != nil {
			return
		}
		err = this.updateFilePath(path, filePathId)
		return
	}
	res, err := this.storage.InsertOne(*this.ctx, bson.M{"path": path})
	if err != nil {
		return
	}
	filePathId = res.InsertedID.(primitive.ObjectID)
	return
}

func (this *FilesService) updateFilePath(path string, filepathId primitive.ObjectID) error {
	var filePath struct {
		Id   primitive.ObjectID
		Path string
	}
	filter := bson.M{"_id": filepathId}
	if err := this.storage.FindOne(*this.ctx, filter).Decode(&filePath); err != nil {
		return err
	}
	if err := os.Remove(filePath.Path); err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{"path": path}}
	_, err := this.storage.UpdateOne(*this.ctx, filter, update)
	return err
}

func (this *FilesService) getFilePath(filename, resourceType string) (path string, err error) {
	ext := filepath.Ext(filename)
	if dir, ok := this.directories[ext]; ok {
		path = fmt.Sprintf("%s/%s/%s/%s", this.basePath, resourceType, dir, filename)
		return
	}
	err = errors.New("This file extension is not supported.")
	return
}
