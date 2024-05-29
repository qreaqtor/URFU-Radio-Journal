package filehand

import (
	"context"
	"io"
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	UploadFile(context.Context, *models.FileUnit) (string, error)
	DownloadFile(context.Context, string) (*models.FileUnit, error)
	DeleteFile(context.Context, string) error
}

type FilesHandler struct {
	files service
}

func NewFilesHandler(files service) *FilesHandler {
	return &FilesHandler{
		files: files,
	}
}

func (fp *FilesHandler) UploadFile(ctx *gin.Context) {
	file, header, err := ctx.Request.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer file.Close()

	fileUnit := &models.FileUnit{
		Payload:     file,
		ContentType: header.Header.Get("Content-Type"),
		Size:        header.Size,
	}

	id, err := fp.files.UploadFile(ctx.Request.Context(), fileUnit)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (fp *FilesHandler) DownloadFile(ctx *gin.Context) {
	fileID := ctx.Param("fileID")

	fileUnit, err := fp.files.DownloadFile(ctx, fileID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	defer fileUnit.Payload.Close()

	_, err = io.Copy(ctx.Writer, fileUnit.Payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	//ctx.Header("Content-Disposition", "attachment")
	ctx.Header("Content-Type", fileUnit.ContentType)
}

func (fp *FilesHandler) DeleteFile(ctx *gin.Context) {
	fileID := ctx.Param("fileID")

	err := fp.files.DeleteFile(ctx.Request.Context(), fileID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
