package articlehand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.ArticleCreate) (string, error)
	GetAll(string) ([]*models.ArticleRead, error)
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
	editionId := ctx.Query("editionId")
	result, err := a.articles.GetAll(editionId)
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

// func (a *ArticleController) deleteContent(articlesFilter, filePathsFilter primitive.M) error {
// 	if err := a.deleteComments(articlesFilter); err != nil {
// 		return err
// 	}
// 	if err := a.deleteFiles(filePathsFilter); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (a *ArticleHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
// 	publicRg.GET("/get/all", a.getAllArticles)
// 	publicRg.GET("/get/:articleId", a.getArticleById)

// 	adminRg.POST("/create", a.create)
// 	adminRg.PUT("/update", a.update)
// 	adminRg.DELETE("/delete/:id", a.delete)
// }

// func (a *ArticleController) GetDeleteHandler() func(primitive.ObjectID) error {
// 	return func(editionId primitive.ObjectID) error {
// 		articlesId, filePathsId, err := a.articles.GetIdsByEditionId(editionId)
// 		articlesFilter := bson.M{"articleId": bson.M{"$in": articlesId}}
// 		filePathsFilter := bson.M{"_id": bson.M{"$in": filePathsId}}
// 		if err != nil {
// 			return err
// 		}
// 		if err := a.deleteContent(articlesFilter, filePathsFilter); err != nil {
// 			return err
// 		}
// 		if err := a.articles.DeleteManyHandler(editionId); err != nil {
// 			return err
// 		}
// 		return nil
// 	}
// }
