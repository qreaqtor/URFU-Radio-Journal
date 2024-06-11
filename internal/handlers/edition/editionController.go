package editionhand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.EditionCreate) (string, error)
	GetAll() ([]*models.EditionRead, error)
	Get(string) (*models.EditionRead, error)
	Update(*models.EditionUpdate) error
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

func (e *EditionHandler) CreateEdition(ctx *gin.Context) {
	edition := &models.EditionCreate{}
	if err := ctx.ShouldBindJSON(edition); err != nil {
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

func (e *EditionHandler) GetAllEditions(ctx *gin.Context) {
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

func (e *EditionHandler) GetEditionById(ctx *gin.Context) {
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

func (e *EditionHandler) UpdateEdition(ctx *gin.Context) {
	edition := &models.EditionUpdate{}
	if err := ctx.ShouldBindJSON(edition); err != nil {
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

func (e *EditionHandler) DeleteEdition(ctx *gin.Context) {
	editionId := ctx.Param("id")

	if err := e.editions.Delete(editionId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}
