package filehand

import (
	"context"
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	UploadFile(context.Context, *models.FileUnit, *models.FileInfo) (string, error)
	DownloadFile(context.Context, string) (*models.FileUnit, error)
	DeleteFile(context.Context, string) error
}

type FilePathsHandler struct {
	files service
}

func NewFilesHandler(files service) *FilePathsHandler {
	return &FilePathsHandler{
		files: files,
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

	fileInfo := &models.FileInfo{
		Filename: fileHeader.Filename,
	}
	fileUnit := &models.FileUnit{
		Payload:     file,
		ContentType: ctx.ContentType(),
		Size:        fileHeader.Size,
	}

	id, err := fp.files.UploadFile(ctx.Request.Context(), fileUnit, fileInfo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (fp *FilePathsHandler) DownloadFile(ctx *gin.Context) {
	fileID := ctx.Param("fileID")

	fileUnit, err := fp.files.DownloadFile(ctx, fileID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	buf := make([]byte, fileUnit.Size)
	_, err = fileUnit.Payload.Read(buf)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.Header("Content-Disposition", "attachment")
	ctx.Data(http.StatusOK, fileUnit.ContentType, buf)
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

func (fp *FilePathsHandler) DeleteFile(ctx *gin.Context) {
	fileID := ctx.Param("fileID")

	err := fp.files.DeleteFile(ctx.Request.Context(), fileID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// func (fp *FilePathsHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
// 	publicRg.GET("/get/:filePathId", fp.getFile)
// 	publicRg.GET("/get/requirements", fp.getRequirements)

// 	adminRg.DELETE("/delete/:filePathId", fp.delete)
// 	adminRg.PUT("/update/:filePathId", fp.updateFile)
// 	adminRg.POST("/upload/:resourceType", fp.uploadFile)
// }
