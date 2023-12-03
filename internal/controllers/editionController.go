package controllers

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditionController struct {
	editions       *services.EditionService
	deleteFiles    func([]primitive.ObjectID) error
	deleteArticles func(primitive.ObjectID) error
}

func NewEditionController(deleteFilesHandler func([]primitive.ObjectID) error, deleteArticlesHandler func(primitive.ObjectID) error) *EditionController {
	return &EditionController{
		editions:       services.NewEditionService(),
		deleteFiles:    deleteFilesHandler,
		deleteArticles: deleteArticlesHandler,
	}
}

func (this *EditionController) createEdition(ctx *gin.Context) {
	var edition models.EditionCreate
	if err := ctx.ShouldBindJSON(&edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := this.editions.Create(edition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *EditionController) getAllEditions(ctx *gin.Context) {
	res, err := this.editions.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (this *EditionController) updateEdition(ctx *gin.Context) {
	var edition models.EditionUpdate
	if err := ctx.ShouldBindJSON(&edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := this.editions.Update(edition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *EditionController) deleteEdition(ctx *gin.Context) {
	editionIdStr := ctx.Param("id")
	edition, err := this.editions.Get(editionIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.deleteContent(edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.editions.Delete(edition.Id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *EditionController) deleteContent(edition models.EditionRead) error {
	if err := this.deleteArticles(edition.Id); err != nil {
		return err
	}
	toDelete := []primitive.ObjectID{edition.CoverPathId, edition.VideoPathId, edition.FilePathId}
	err := this.deleteFiles(toDelete)
	return err
}

func (this *EditionController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAllEditions)

	adminRg.POST("/create", this.createEdition)
	adminRg.PUT("/update", this.updateEdition)
	adminRg.DELETE("/delete/:id", this.deleteEdition)
}
