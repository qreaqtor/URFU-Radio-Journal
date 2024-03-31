package edition

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services/edition"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EditionController struct {
	editions       *edition.EditionService
	deleteFiles    func(filter primitive.M) error
	deleteArticles func(primitive.ObjectID) error
}

func NewEditionController(deleteFilesHandler func(filter primitive.M) error, deleteArticlesHandler func(primitive.ObjectID) error) *EditionController {
	return &EditionController{
		editions:       edition.NewEditionService(),
		deleteFiles:    deleteFilesHandler,
		deleteArticles: deleteArticlesHandler,
	}
}

func (e *EditionController) createEdition(ctx *gin.Context) {
	var edition models.EditionCreate
	if err := ctx.ShouldBindJSON(&edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	id, err := e.editions.Create(edition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"id":      id,
	})
}

func (e *EditionController) getAllEditions(ctx *gin.Context) {
	res, err := e.editions.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    res,
		"message": "success",
	})
}

func (e *EditionController) getEditionById(ctx *gin.Context) {
	editionId := ctx.Param("editionId")
	edition, err := e.editions.Get(editionId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"edition": edition,
		"message": "success",
	})
}

func (e *EditionController) updateEdition(ctx *gin.Context) {
	var edition models.EditionUpdate
	if err := ctx.ShouldBindJSON(&edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := e.editions.Update(edition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (e *EditionController) deleteEdition(ctx *gin.Context) {
	editionIdStr := ctx.Param("id")
	edition, err := e.editions.Get(editionIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := e.deleteContent(edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := e.editions.Delete(edition.Id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (e *EditionController) deleteContent(edition models.EditionRead) error {
	if err := e.deleteArticles(edition.Id); err != nil {
		return err
	}
	toDelete := []primitive.ObjectID{edition.CoverPathId, edition.VideoPathId, edition.FilePathId}
	filter := bson.M{"_id": bson.M{"$in": toDelete}}
	err := e.deleteFiles(filter)
	return err
}

func (e *EditionController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", e.getAllEditions)
	publicRg.GET("/get/:editionId", e.getEditionById)

	adminRg.POST("/create", e.createEdition)
	adminRg.PUT("/update", e.updateEdition)
	adminRg.DELETE("/delete/:id", e.deleteEdition)
}
