package controllers

import (
	"net/http"
	"strconv"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilePathsController struct {
	filePaths *services.FilePathsService
}

func NewFilesController() *FilePathsController {
	return &FilePathsController{
		filePaths: services.NewFilesService(),
	}
}

func (this *FilePathsController) uploadFile(ctx *gin.Context) {
	resourceType := ctx.Param("resourceType")
	err := this.filePaths.CheckResourceType(resourceType)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, path, err := this.filePaths.GetFilePathInfo(file.Filename, resourceType)
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

func (this *FilePathsController) updateFile(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePathId := ctx.Param("filePathId")
	path, err := this.filePaths.UpdateFile(file.Filename, filePathId)
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

func (this *FilePathsController) getFile(ctx *gin.Context) {
	download, err := strconv.ParseBool(ctx.Query("download"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if download {
		ctx.Header("Content-Disposition", "attachment")
	}
	filePathId := ctx.Param("filePathId")
	path, err := this.filePaths.CheckFilePath(filePathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.File(path)
}

func (this *FilePathsController) delete(ctx *gin.Context) {
	filePathId := ctx.Param("filePathId")
	err := this.filePaths.DeleteOne(filePathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *FilePathsController) getRequirements(ctx *gin.Context) {
	data, err := this.filePaths.GetRequirementsFiles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"data":    data,
	})
}

func (this *FilePathsController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/:filePathId", this.getFile)
	publicRg.GET("/get/requirements", this.getRequirements)

	adminRg.DELETE("/delete/:filePathId", this.delete)
	adminRg.PUT("/update/:filePathId", this.updateFile)
	adminRg.POST("/upload/:resourceType", this.uploadFile)
}

func (this *FilePathsController) GetDeleteHandler() func(filter primitive.M) error {
	return this.filePaths.DeleteManyHandler
}
