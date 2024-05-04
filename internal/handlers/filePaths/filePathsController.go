package filepathand

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type service interface {
	DeleteOne(string) error
	GetRequirementsFiles() ([]string, error)
	CheckResourceType(string) error
	CheckFilePath(string) (string, error)
	GetFilePathInfo(string, string) (string, string, error)
	UpdateFile(string, string) (string, error)
}

type FilePathsHandler struct {
	filePaths service
}

func NewFileshandler(filePaths service) *FilePathsHandler {
	return &FilePathsHandler{
		filePaths: filePaths,
	}
}

func (fp *FilePathsHandler) uploadFile(ctx *gin.Context) {
	resourceType := ctx.Param("resourceType")
	err := fp.filePaths.CheckResourceType(resourceType)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, path, err := fp.filePaths.GetFilePathInfo(file.Filename, resourceType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (fp *FilePathsHandler) updateFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePathId := ctx.Param("filePathId")
	path, err := fp.filePaths.UpdateFile(file.Filename, filePathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (fp *FilePathsHandler) getFile(ctx *gin.Context) {
	downloadStr := ctx.Query("download")
	if downloadStr != "" {
		download, err := strconv.ParseBool(downloadStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if download {
			ctx.Header("Content-Disposition", "attachment")
		}
	}
	filePathId := ctx.Param("filePathId")
	path, err := fp.filePaths.CheckFilePath(filePathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.File(path)
}

func (fp *FilePathsHandler) delete(ctx *gin.Context) {
	filePathId := ctx.Param("filePathId")
	err := fp.filePaths.DeleteOne(filePathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (fp *FilePathsHandler) getRequirements(ctx *gin.Context) {
	data, err := fp.filePaths.GetRequirementsFiles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (fp *FilePathsHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/:filePathId", fp.getFile)
	publicRg.GET("/get/requirements", fp.getRequirements)

	adminRg.DELETE("/delete/:filePathId", fp.delete)
	adminRg.PUT("/update/:filePathId", fp.updateFile)
	adminRg.POST("/upload/:resourceType", fp.uploadFile)
}

// func (fp *FilePathsHandler) GetDeleteHandler() func(filter primitive.M) error {
// 	return fp.filePaths.DeleteManyHandler
// }
