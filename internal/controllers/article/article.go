package article

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services/article"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleController struct {
	articles       *article.ArticleService
	deleteFiles    func(filter primitive.M) error
	deleteComments func(filter primitive.M) error
}

func NewArticleController(deleteFileHandler func(filter primitive.M) error, deleteCommentsHandler func(filter primitive.M) error) *ArticleController {
	return &ArticleController{
		articles:       article.NewArticleService(),
		deleteFiles:    deleteFileHandler,
		deleteComments: deleteCommentsHandler,
	}
}

func (a *ArticleController) create(ctx *gin.Context) {
	var article models.ArticleCreate
	if err := ctx.ShouldBindJSON(&article); err != nil {
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

func (a *ArticleController) getAllArticles(ctx *gin.Context) {
	editionId := ctx.Query("editionId")
	result, err := a.articles.GetAll(editionId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": result})
}

func (a *ArticleController) getArticleById(ctx *gin.Context) {
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

func (a *ArticleController) update(ctx *gin.Context) {
	var article models.ArticleUpdate
	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := a.articles.Update(article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (a *ArticleController) delete(ctx *gin.Context) {
	articleIdStr := ctx.Param("id")
	articleId, err := primitive.ObjectIDFromHex(articleIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filePathId, err := a.articles.GetFilePathId(articleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	articlesFilter := bson.M{"articleId": articleId}
	filePathsFilter := bson.M{"_id": filePathId}
	if err := a.deleteContent(articlesFilter, filePathsFilter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := a.articles.Delete(articleId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (a *ArticleController) deleteContent(articlesFilter, filePathsFilter primitive.M) error {
	if err := a.deleteComments(articlesFilter); err != nil {
		return err
	}
	if err := a.deleteFiles(filePathsFilter); err != nil {
		return err
	}
	return nil
}

func (a *ArticleController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", a.getAllArticles)
	publicRg.GET("/get/:articleId", a.getArticleById)

	adminRg.POST("/create", a.create)
	adminRg.PUT("/update", a.update)
	adminRg.DELETE("/delete/:id", a.delete)
}

func (a *ArticleController) GetDeleteHandler() func(primitive.ObjectID) error {
	return func(editionId primitive.ObjectID) error {
		articlesId, filePathsId, err := a.articles.GetIdsByEditionId(editionId)
		articlesFilter := bson.M{"articleId": bson.M{"$in": articlesId}}
		filePathsFilter := bson.M{"_id": bson.M{"$in": filePathsId}}
		if err != nil {
			return err
		}
		if err := a.deleteContent(articlesFilter, filePathsFilter); err != nil {
			return err
		}
		if err := a.articles.DeleteManyHandler(editionId); err != nil {
			return err
		}
		return nil
	}
}
