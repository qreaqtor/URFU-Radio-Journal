package filePaths

import (
	"net/http"
	"strconv"
	"urfu-radio-journal/pkg/services/filePaths"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilePathsController struct {
	filePaths *filePaths.FilePathsService
}

func NewFilesController() *FilePathsController {
	return &FilePathsController{
		filePaths: filePaths.NewFilesService(),
	}
}

func (fp *FilePathsController) uploadFile(ctx *gin.Context) {
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

func (fp *FilePathsController) updateFile(ctx *gin.Context) {
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

func (fp *FilePathsController) getFile(ctx *gin.Context) {
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

func (fp *FilePathsController) delete(ctx *gin.Context) {
	filePathId := ctx.Param("filePathId")
	err := fp.filePaths.DeleteOne(filePathId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (fp *FilePathsController) getRequirements(ctx *gin.Context) {
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

func (fp *FilePathsController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/:filePathId", fp.getFile)
	publicRg.GET("/get/requirements", fp.getRequirements)

	adminRg.DELETE("/delete/:filePathId", fp.delete)
	adminRg.PUT("/update/:filePathId", fp.updateFile)
	adminRg.POST("/upload/:resourceType", fp.uploadFile)
}

func (fp *FilePathsController) GetDeleteHandler() func(filter primitive.M) error {
	return fp.filePaths.DeleteManyHandler
}
