package filehand

import (
	"context"
	"fmt"
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

type monitoring interface {
	UpdateDownloads(string)
}

type FilesHandler struct {
	files service
	downloads monitoring
}

func NewFilesHandler(files service, downloads monitoring) *FilesHandler {
	return &FilesHandler{
		files: files,
		downloads: downloads,
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
		Filename:    header.Filename,
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

	fp.downloads.UpdateDownloads(fileUnit.Filename)

	ctx.Header("Content-Disposition", fmt.Sprintf(`inline; filename="%s"`, fileUnit.Filename))
	ctx.Header("Content-Type", fileUnit.ContentType)

	_, err = io.Copy(ctx.Writer, fileUnit.Payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
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
