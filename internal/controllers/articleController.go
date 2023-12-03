package controllers

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleController struct {
	articles       *services.ArticleService
	deleteFile     func([]primitive.ObjectID) error
	deleteComments func([]primitive.ObjectID) error
}

func NewArticleController(deleteFileHandler func([]primitive.ObjectID) error, deleteCommentsHandler func([]primitive.ObjectID) error) *ArticleController {
	return &ArticleController{
		articles:       services.NewArticleService(),
		deleteFile:     deleteFileHandler,
		deleteComments: deleteCommentsHandler,
	}
}

func (this *ArticleController) create(ctx *gin.Context) {
	var article models.ArticleCreate
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	articleId, err := this.articles.Create(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      articleId,
	})
}

func (this *ArticleController) getAllArticles(ctx *gin.Context) {
	result, err := this.articles.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (this *ArticleController) update(ctx *gin.Context) {
	var article models.ArticleUpdate
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.articles.Update(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *ArticleController) delete(ctx *gin.Context) {
	// var input struct {
	// 	Data []primitive.ObjectID `json:"data" binding:"required"`
	// }
	// if err := ctx.ShouldBindJSON(&input); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	articleIdStr := ctx.Param("id")
	articleId, err := primitive.ObjectIDFromHex(articleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePathId, err := this.articles.GetFilePathId(articleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.deleteContent([]primitive.ObjectID{articleId}, []primitive.ObjectID{filePathId}); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.articles.Delete(articleId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *ArticleController) deleteContent(articlesId, filePathsId []primitive.ObjectID) error {
	if err := this.deleteComments(articlesId); err != nil {
		return err
	}
	if err := this.deleteFile(filePathsId); err != nil {
		return err
	}
	return nil
}

func (this *ArticleController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAllArticles)

	adminRg.POST("/create", this.create)
	adminRg.PUT("/update", this.update)
	adminRg.DELETE("/delete/:id", this.delete)
}

func (this *ArticleController) GetDeleteHandler() func(primitive.ObjectID) error {
	return func(editionId primitive.ObjectID) error {
		articlesId, filePathsId, err := this.articles.GetIdsByEditionId(editionId)
		if err != nil {
			return err
		}
		if err := this.deleteContent(articlesId, filePathsId); err != nil {
			return err
		}
		if err := this.articles.DeleteMany(editionId); err != nil {
			return err
		}
		return nil
	}
}
