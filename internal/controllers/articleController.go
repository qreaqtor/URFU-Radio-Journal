package controllers

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	articles *services.ArticleService
}

func NewArticlecontroller() *ArticleController {
	return &ArticleController{articles: services.NewArticleService()}
}

func (this *ArticleController) createArticle(ctx *gin.Context) {
	var article models.ArticleCreate
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.articles.CreateArticle(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *ArticleController) getAllArticles(ctx *gin.Context) {
	result, err := this.articles.GetAllArticles()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (this *ArticleController) updateArticle(ctx *gin.Context) {
	var article models.ArticleUpdate
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.articles.UpdateArticle(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *ArticleController) deleteArticle(ctx *gin.Context) {
	id := ctx.Query("id")
	if err := this.articles.DeleteArticle(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *ArticleController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAllArticles)

	adminRg.POST("/create", this.createArticle)
	adminRg.PUT("/update", this.updateArticle)
	adminRg.DELETE("/delete/:id", this.deleteArticle)
}
