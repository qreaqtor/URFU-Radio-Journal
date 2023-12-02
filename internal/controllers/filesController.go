package controllers

import (
	"net/http"
	"strings"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
)

type FilesController struct {
	files *services.FilesService
}

func NewFilesController() *FilesController {
	return &FilesController{files: services.NewFilesService()}
}

func (this *FilesController) uploadFile(ctx *gin.Context) {
	resourceType := ctx.MustGet("resourceType").(string)
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePathId := ctx.Query("filePathId")
	url, path, err := this.files.GetFileURL(file.Filename, resourceType, filePathId)
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
		"url":     url,
	})
}

func (this *FilesController) downloadFile(ctx *gin.Context) {
	resourceType := ctx.MustGet("resourceType").(string)
	filePathId := ctx.Param("filePathId")
	path, err := this.files.CheckFilePath(filePathId, resourceType)
	if err == nil {
		ctx.File(path)
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

func (this *FilesController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.Use(this.resourceTypeMiddleware())
	adminRg.Use(this.resourceTypeMiddleware())

	publicRg.GET("/editions/download/:filePathId", this.downloadFile)
	publicRg.GET("/articles/download/:filePathId", this.downloadFile)

	adminRg.POST("/editions/upload", this.uploadFile)
	adminRg.POST("/articles/upload", this.uploadFile)
}

func (this *FilesController) resourceTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqPath := ctx.Request.URL.Path
		pathParts := strings.Split(reqPath, "/")
		if pathParts[3] != "editions" && pathParts[3] != "articles" {
			ctx.JSON(http.StatusBadGateway, gin.H{"message": "Wrong gateway."})
			ctx.Abort()
			return
		}
		ctx.Set("resourceType", pathParts[3])
		ctx.Next()
		return
	}
}
