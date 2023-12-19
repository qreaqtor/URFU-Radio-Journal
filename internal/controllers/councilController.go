package controllers

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CouncilController struct {
	members     *services.CouncilService
	deleteFiles func(filter primitive.M) error
}

func NewCouncilController(deleteFilesHandler func(filter primitive.M) error) *CouncilController {
	return &CouncilController{
		members:     services.NewCouncilService(),
		deleteFiles: deleteFilesHandler,
	}
}

func (this *CouncilController) create(ctx *gin.Context) {
	var member models.CouncilMemberCreate
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.members.Create(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CouncilController) update(ctx *gin.Context) {
	var member models.CouncilMemberUpdate
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	memberId := ctx.Param("id")
	if err := this.members.Update(memberId, member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CouncilController) getAll(ctx *gin.Context) {
	data, err := this.members.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": data})
}

func (this *CouncilController) delete(ctx *gin.Context) {
	memberIdStr := ctx.Param("id")
	memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	imagePathId, err := this.members.GetImagePathId(memberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filter := bson.M{"_id": imagePathId}
	if err := this.deleteFiles(filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.members.Delete(memberId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CouncilController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAll)

	adminRg.POST("/create", this.create)
	adminRg.PUT("/update/:id", this.update)
	adminRg.DELETE("/delete/:id", this.delete)
}
