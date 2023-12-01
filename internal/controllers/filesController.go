package controllers

import (
	"net/http"
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
	resourceType := ctx.Param("resourceType")
	if resourceType != "editions" && resourceType != "articles" {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Wrong gateway."})
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePath, err := this.files.GetFilePath(file.Filename, resourceType)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	identifier := ctx.Query("identifier")
	url, err := this.files.GetFileURL(file.Filename, resourceType, filePath, identifier)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = ctx.SaveUploadedFile(file, filePath)
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
	resourceType := ctx.Param("resourceType")
	if resourceType != "editions" && resourceType != "articles" {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": "Wrong gateway."})
		return
	}
	filename := ctx.Param("filename")
	path, err := this.files.CheckFilePath(filename, resourceType)
	if err == nil {
		ctx.File(path)
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
}

func (this *FilesController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.GET("/download/:resourceType/:filename", this.downloadFile)

	adminRg.POST("/upload/:resourceType", this.uploadFile)
}
