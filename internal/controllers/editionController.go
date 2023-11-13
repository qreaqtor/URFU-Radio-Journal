package controllers

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
)

type EditionController struct {
	service *services.EditionService
}

func NewEditionController() *EditionController {
	return &EditionController{service: services.NewEditionService()}
}

func (this *EditionController) createEdition(ctx *gin.Context) {
	var edition models.Edition
	if err := ctx.ShouldBindJSON(&edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := this.service.CreateEdition(edition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *EditionController) getAllEditions(ctx *gin.Context) {
	res, err := this.service.GetAllEditions()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": res})
}

func (this *EditionController) updateEdition(ctx *gin.Context) {
	var edition models.Edition
	if err := ctx.ShouldBindJSON(&edition); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := this.service.UpdateEdition(edition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *EditionController) deleteEdition(ctx *gin.Context) {
	editionId := ctx.Param("id")
	if err := this.service.DeleteEdition(editionId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *EditionController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/create", this.createEdition)
	rg.GET("/get/all", this.getAllEditions)
	rg.PUT("/update", this.updateEdition)
	rg.DELETE("/delete/:id", this.deleteEdition)
}
