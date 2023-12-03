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

type FilePathsService struct {
	ctx         *context.Context
	storage     *mongo.Collection
	basePath    string
	directories map[string]string
}

func NewFilesService() *FilePathsService {
	return &FilePathsService{
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

func (this *FilePathsService) CheckFilePath(filePathIdStr string) (path string, err error) {
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

func (this *FilePathsService) GetFileURL(filename, resourceType, filePathIdStr string) (filePathId primitive.ObjectID, path string, err error) {
	if path, err = this.generateFilePath(filename, resourceType); err != nil {
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

func (this *FilePathsService) DeleteHandler(data []primitive.ObjectID) error {
	var filePaths []struct {
		Id   primitive.ObjectID
		Path string
	}
	filter := bson.M{"_id": bson.M{"$in": data}}
	cur, err := this.storage.Find(*this.ctx, filter)
	if err := cur.All(*this.ctx, &filePaths); err != nil {
		return err
	}
	for _, v := range filePaths {
		path := v.Path
		if err := os.Remove(path); err != nil {
			return err
		}
	}
	res, err := this.storage.DeleteMany(*this.ctx, filter)
	if res.DeletedCount == 0 {
		return errors.New("Document not found.")
	}
	return err
}

func (this *FilePathsService) updateFilePath(path string, filepathId primitive.ObjectID) error {
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

func (this *FilePathsService) generateFilePath(filename, resourceType string) (path string, err error) {
	ext := filepath.Ext(filename)
	if dir, ok := this.directories[ext]; ok {
		path = fmt.Sprintf("%s/%s/%s/%s", this.basePath, resourceType, dir, filename)
		return
	}
	err = errors.New("This file extension is not supported.")
	return
}
