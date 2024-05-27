package filepathand

import (
	"context"
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(context.Context, *models.FileUnit) (string, error)
	Get(context.Context) error
	Update(context.Context) error
	Delete(context.Context, string) error
}

type FilePathsHandler struct {
	filePaths service
}

func NewFileshandler(filePaths service) *FilePathsHandler {
	return &FilePathsHandler{
		filePaths: filePaths,
	}
}

func (fp *FilePathsHandler) UploadFile(ctx *gin.Context) {
	fileHeader, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	fileUnit := &models.FileUnit{
		Payload: file,
		Info: &models.FileInfo{
			Name:        fileHeader.Filename,
			ContentType: ctx.ContentType(),
			Size:        fileHeader.Size,
		},
	}

	id, err := fp.filePaths.Create(ctx.Request.Context(), fileUnit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (fp *FilePathsHandler) UpdateFile(ctx *gin.Context) {
	// file, err := ctx.FormFile("file")
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// filePathId := ctx.Param("filePathId")
	// path, err := fp.filePaths.UpdateFile(file.Filename, filePathId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// err = ctx.SaveUploadedFile(file, path)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (fp *FilePathsHandler) Get(ctx *gin.Context) {
	// downloadStr := ctx.Query("download")
	// if downloadStr != "" {
	// 	download, err := strconv.ParseBool(downloadStr)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// 	if download {
	// 		ctx.Header("Content-Disposition", "attachment")
	// 	}
	// }
	// filePathId := ctx.Param("filePathId")
	// path, err := fp.filePaths.CheckFilePath(filePathId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// ctx.File(path)
}

func (fp *FilePathsHandler) Delete(ctx *gin.Context) {
	// filePathId := ctx.Param("filePathId")
	// err := fp.filePaths.DeleteOne(filePathId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// func (fp *FilePathsHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
// 	publicRg.GET("/get/:filePathId", fp.getFile)
// 	publicRg.GET("/get/requirements", fp.getRequirements)

// 	adminRg.DELETE("/delete/:filePathId", fp.delete)
// 	adminRg.PUT("/update/:filePathId", fp.updateFile)
// 	adminRg.POST("/upload/:resourceType", fp.uploadFile)
// }

// func (fp *FilePathsHandler) GetDeleteHandler() func(filter primitive.M) error {
// 	return fp.filePaths.DeleteManyHandler
// }
