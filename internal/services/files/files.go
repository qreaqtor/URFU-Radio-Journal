package filesrv

import (
	"context"
	"errors"
	"fmt"
	"urfu-radio-journal/internal/models"
)

var (
	errBadContentType = errors.New("bad Content-Type")
)

type UploadFileFunc func(context.Context, *models.FileUnit) error

type fileRepo interface {
	UploadFile(context.Context, *models.FileUnit) error
	DeleteFile(context.Context, string) error
}

type fileInfo interface {
	InsertOne(*models.FileInfo) (string, error)
	DeleteOne(string) error
	FindOne(string) (*models.FileInfo, error)
	UpdateOne(string) error
}

type FileService struct {
	videos    fileRepo
	documents fileRepo
	images    fileRepo
	files     fileInfo
}

func NewFileService(videos, documents, images fileRepo, files fileInfo) *FileService {
	return &FileService{
		videos:    videos,
		documents: documents,
		images:    images,
		files:     files,
	}
}

func (f *FileService) getUploadFunc(fileContentType string) (UploadFileFunc, error) {
	switch fileContentType {
	case "video/mp4":
		return f.videos.UploadFile, nil
	case "image/jpeg":
		return f.images.UploadFile, nil
	case "application/pdf":
		return f.documents.UploadFile, nil
	}
	return nil, errBadContentType
}

func (f *FileService) Create(ctx context.Context, file *models.FileUnit) (string, error) {
	upload, err := f.getUploadFunc(file.Info.ContentType)
	if err != nil {
		return "", err
	}

	id, err := f.files.InsertOne(file.Info)
	if err != nil {
		return "", err
	}

	file.PayloadID = id

	err = upload(ctx, file)
	if err != nil {
		err2 := f.files.DeleteOne(id)
		if err2 != nil {
			return "", fmt.Errorf("%v; %v", err, err2)
		}
		return "", err
	}

	return file.PayloadID, nil
}

func (f *FileService) Get(ctx context.Context, id string) error {
	return nil
}

func (f *FileService) Update(ctx context.Context, id string) error {
	return nil
}

func (f *FileService) Delete(ctx context.Context, fileName string) error {
	err := f.videos.DeleteFile(ctx, fileName)
	return err
}

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"os"
// 	"path/filepath"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type FilePathsService struct {
// 	storage       *mongo.Collection
// 	basePath      string
// 	directories   map[string]string
// 	resourceTypes [4]string
// }

// func NewFilesService() *FilePathsService {
// 	return &FilePathsService{
// 		ctx:           db.GetContext(),
// 		storage:       db.GetStorage("filePaths"),
// 		basePath:      "../attachments",
// 		directories:   getDirs(),
// 		resourceTypes: getResourceTypes(),
// 	}
// }

// func getDirs() map[string]string {
// 	dirs := make(map[string]string, 4)
// 	dirs[".pdf"] = "documents"
// 	dirs[".jpeg"] = "images"
// 	dirs[".jpg"] = "images"
// 	dirs[".mkv"] = "videos"
// 	dirs[".mp4"] = "videos"
// 	return dirs
// }

// func getResourceTypes() [4]string {
// 	resourceTypes := [4]string{"requirements", "editions", "articles", "avatars"}
// 	return resourceTypes
// }

// func (fps *FilePathsService) CheckFilePath(filePathIdStr string) (path string, err error) {
// 	filePathId, err := primitive.ObjectIDFromHex(filePathIdStr)
// 	if err != nil {
// 		return
// 	}
// 	var filePath struct {
// 		Id   primitive.ObjectID
// 		Path string
// 	}
// 	filter := bson.M{"_id": filePathId}
// 	err = fps.storage.FindOne(*fps.ctx, filter).Decode(&filePath)
// 	path = filePath.Path
// 	return
// }

// func (fps *FilePathsService) GetFilePathInfo(filename, resourceType string) (filePathId primitive.ObjectID, path string, err error) {
// 	if path, err = fps.generateFilePath(filename); err != nil {
// 		return
// 	}
// 	file := bson.M{
// 		"path":         path,
// 		"resourceType": resourceType,
// 	}
// 	res, err := fps.storage.InsertOne(*fps.ctx, file)
// 	if err != nil {
// 		return
// 	}
// 	filePathId = res.InsertedID.(primitive.ObjectID)
// 	return
// }

// func (fps *FilePathsService) GetRequirementsFiles() (files map[string]primitive.ObjectID, err error) {
// 	var filePaths []struct {
// 		Id   primitive.ObjectID `bson:"_id"`
// 		Path string             `bson:"path"`
// 	}
// 	filter := bson.M{"resourceType": "requirements"}
// 	cur, err := fps.storage.Find(*fps.ctx, filter)
// 	if err != nil {
// 		return
// 	}
// 	if err = cur.All(*fps.ctx, &filePaths); err != nil {
// 		return
// 	}
// 	files = make(map[string]primitive.ObjectID, len(filePaths))
// 	for _, v := range filePaths {
// 		fileName := filepath.Base(v.Path)
// 		files[fileName] = v.Id
// 	}
// 	return
// }

// func (fps *FilePathsService) DeleteOne(idStr string) error {
// 	id, err := primitive.ObjectIDFromHex(idStr)
// 	if err != nil {
// 		return err
// 	}
// 	var filePath struct {
// 		Id   primitive.ObjectID
// 		Path string
// 	}
// 	filter := bson.M{"_id": id}
// 	err = fps.storage.FindOneAndDelete(*fps.ctx, filter).Decode(&filePath)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Remove(filePath.Path)
// 	return err
// }

// func (fps *FilePathsService) DeleteManyHandler(filter primitive.M) error {
// 	var filePaths []struct {
// 		Id   primitive.ObjectID
// 		Path string
// 	}
// 	cur, err := fps.storage.Find(*fps.ctx, filter)
// 	if err != nil {
// 		return err
// 	}
// 	if err = cur.All(*fps.ctx, &filePaths); err != nil {
// 		return err
// 	}
// 	for _, v := range filePaths {
// 		path := v.Path
// 		if err := os.Remove(path); err != nil {
// 			return err
// 		}
// 	}
// 	_, err = fps.storage.DeleteMany(*fps.ctx, filter)
// 	return err
// }

// func (fps *FilePathsService) UpdateFile(filename, filePathIdStr string) (path string, err error) {
// 	path, err = fps.generateFilePath(filename)
// 	if err != nil {
// 		return
// 	}
// 	filePathId, err := primitive.ObjectIDFromHex(filePathIdStr)
// 	if err != nil {
// 		return
// 	}
// 	err = fps.updateFilePath(path, filePathId)
// 	return
// }

// func (fps *FilePathsService) CheckResourceType(resourceType string) error {
// 	for _, v := range fps.resourceTypes {
// 		if resourceType == v {
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("incorrect resource type: %s", resourceType)
// }

// func (fps *FilePathsService) updateFilePath(path string, filepathId primitive.ObjectID) error {
// 	var filePath struct {
// 		Path string `bson:"path"`
// 	}
// 	filter := bson.M{"_id": filepathId}
// 	update := bson.M{"$set": bson.M{"path": path}}
// 	returnDoc := options.Before
// 	options := options.FindOneAndUpdateOptions{
// 		ReturnDocument: &returnDoc,
// 	}
// 	err := fps.storage.FindOneAndUpdate(*fps.ctx, filter, update, &options).Decode(&filePath)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.Remove(filePath.Path)
// 	return err
// }

// func (fps *FilePathsService) generateFilePath(filename string) (path string, err error) {
// 	ext := filepath.Ext(filename)
// 	if dir, ok := fps.directories[ext]; ok {
// 		path = fmt.Sprintf("%s/%s/%s", fps.basePath, dir, filename)
// 		return
// 	}
// 	err = errors.New("this file extension is not supported")
// 	return
// }
