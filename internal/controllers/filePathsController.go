package controllers

import (
	"net/http"
	"strings"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilePathsController struct {
	filePaths *services.FilePathsService
}

func NewFilesController() *FilePathsController {
	return &FilePathsController{filePaths: services.NewFilesService()}
}

func (this *FilePathsController) uploadFile(ctx *gin.Context) {
	resourceType := ctx.MustGet("resourceType").(string)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePathId := ctx.Query("filePathId")
	id, path, err := this.filePaths.GetFileURL(file.Filename, resourceType, filePathId)
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

func (this *FilePathsController) downloadFile(ctx *gin.Context) {
	filePathId := ctx.Param("filePathId")
	path, err := this.filePaths.CheckFilePath(filePathId)
	if err == nil {
		ctx.File(path)
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

func (this *FilePathsController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	adminRg.Use(this.resourceTypeMiddleware())

	publicRg.GET("/download/:filePathId", this.downloadFile)

	adminRg.POST("/editions/upload", this.uploadFile)
	adminRg.POST("/articles/upload", this.uploadFile)
}

func (this *FilePathsController) resourceTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqPath := ctx.Request.URL.Path
		pathParts := strings.Split(reqPath, "/")
		ctx.Set("resourceType", pathParts[3])
		ctx.Next()
		return
	}
}

func (this *FilePathsController) GetDeleteHandler() func(filter primitive.M) error {
	return this.filePaths.DeleteManyHandler
}
