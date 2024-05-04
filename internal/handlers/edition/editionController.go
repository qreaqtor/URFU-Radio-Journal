package editionhand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(models.EditionCreate) (string, error)
	GetAll() ([]models.EditionRead, error)
	Get(string) (models.EditionRead, error)
	Update(models.EditionUpdate) error
	Delete(string) error
}

type EditionHandler struct {
	editions service
}

func NewEditionHandler(edition service) *EditionHandler {
	return &EditionHandler{
		editions: edition,
	}
}

func (e *EditionHandler) createEdition(ctx *gin.Context) {
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

func (e *EditionHandler) getAllEditions(ctx *gin.Context) {
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

func (e *EditionHandler) getEditionById(ctx *gin.Context) {
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

func (e *EditionHandler) updateEdition(ctx *gin.Context) {
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

func (e *EditionHandler) deleteEdition(ctx *gin.Context) {
	editionId := ctx.Param("id")
	// edition, err := e.editions.Get(editionIdStr)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// if err := e.deleteContent(edition); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	if err := e.editions.Delete(editionId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// func (e *EditionHandler) deleteContent(edition models.EditionRead) error {
// 	if err := e.deleteArticles(edition.Id); err != nil {
// 		return err
// 	}
// 	toDelete := []primitive.ObjectID{edition.CoverPathId, edition.VideoPathId, edition.FilePathId}
// 	filter := bson.M{"_id": bson.M{"$in": toDelete}}
// 	err := e.deleteFiles(filter)
// 	return err
// }

func (e *EditionHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", e.getAllEditions)
	publicRg.GET("/get/:editionId", e.getEditionById)

	adminRg.POST("/create", e.createEdition)
	adminRg.PUT("/update", e.updateEdition)
	adminRg.DELETE("/delete/:id", e.deleteEdition)
}
