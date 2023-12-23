package controllers

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleController struct {
	articles       *services.ArticleService
	deleteFiles    func(filter primitive.M) error
	deleteComments func(filter primitive.M) error
}

func NewArticleController(deleteFileHandler func(filter primitive.M) error, deleteCommentsHandler func(filter primitive.M) error) *ArticleController {
	return &ArticleController{
		articles:       services.NewArticleService(),
		deleteFiles:    deleteFileHandler,
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

func (this *ArticleController) getArticlesByEdition(ctx *gin.Context) {
	editionId := ctx.Param(":editionId")
	result, err := this.articles.GetAllByEditionId(editionId)
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
	articlesFilter := bson.M{"articleId": articleId}
	filePathsFilter := bson.M{"_id": filePathId}
	if err := this.deleteContent(articlesFilter, filePathsFilter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.articles.Delete(articleId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *ArticleController) deleteContent(articlesFilter, filePathsFilter primitive.M) error {
	if err := this.deleteComments(articlesFilter); err != nil {
		return err
	}
	if err := this.deleteFiles(filePathsFilter); err != nil {
		return err
	}
	return nil
}

func (this *ArticleController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAllArticles)
	publicRg.GET("/get/:editionId", this.getArticlesByEdition)

	adminRg.POST("/create", this.create)
	adminRg.PUT("/update", this.update)
	adminRg.DELETE("/delete/:id", this.delete)
}

func (this *ArticleController) GetDeleteHandler() func(primitive.ObjectID) error {
	return func(editionId primitive.ObjectID) error {
		articlesId, filePathsId, err := this.articles.GetIdsByEditionId(editionId)
		articlesFilter := bson.M{"articleId": bson.M{"$in": articlesId}}
		filePathsFilter := bson.M{"_id": bson.M{"$in": filePathsId}}
		if err != nil {
			return err
		}
		if err := this.deleteContent(articlesFilter, filePathsFilter); err != nil {
			return err
		}
		if err := this.articles.DeleteManyHandler(editionId); err != nil {
			return err
		}
		return nil
	}
}
