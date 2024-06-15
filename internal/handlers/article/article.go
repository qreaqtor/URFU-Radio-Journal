package articlehand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.ArticleCreate) (string, error)
	GetAll(*models.ArticleQuery) ([]*models.ArticleRead, error)
	Get(string) (*models.ArticleRead, error)
	Update(*models.ArticleUpdate) error
	Delete(string) error
}

type ArticleHandler struct {
	articles service
}

func NewArticleHandler(articles service) *ArticleHandler {
	return &ArticleHandler{
		articles: articles,
	}
}

func (a *ArticleHandler) Create(ctx *gin.Context) {
	article := &models.ArticleCreate{}
	if err := ctx.ShouldBindJSON(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	articleId, err := a.articles.Create(article)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      articleId,
	})
}

func (a *ArticleHandler) GetAllArticles(ctx *gin.Context) {
	args := &models.ArticleQuery{}
	err := ctx.ShouldBindQuery(args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	result, err := a.articles.GetAll(args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (a *ArticleHandler) GetArticleById(ctx *gin.Context) {
	articleId := ctx.Param("articleId")
	article, err := a.articles.Get(articleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"article": article,
		"message": "success",
	})
}

func (a *ArticleHandler) Update(ctx *gin.Context) {
	article := &models.ArticleUpdate{}
	if err := ctx.ShouldBindJSON(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := a.articles.Update(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (a *ArticleHandler) Delete(ctx *gin.Context) {
	articleId := ctx.Param("id")

	if err := a.articles.Delete(articleId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
