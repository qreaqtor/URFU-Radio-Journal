package controllers

import (
	"net/http"
	"strconv"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentsController struct {
	comments *services.CommentsService
}

func NewCommentsController() *CommentsController {
	return &CommentsController{comments: services.NewCommentsService()}
}

func (this *CommentsController) create(ctx *gin.Context) {
	var comment models.CommentCreate
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Create(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) getAll(ctx *gin.Context) {
	onlyApproved, err := strconv.ParseBool(ctx.Query("onlyApproved"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	comments, err := this.comments.GetAll(onlyApproved)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": comments})
}

func (this *CommentsController) update(ctx *gin.Context) {
	var input struct {
		Data []models.CommentUpdate `json:"data" binding:"required,dive"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Update(input.Data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) delete(ctx *gin.Context) {
	var input struct {
		Data []primitive.ObjectID `json:"data" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Delete(input.Data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) approve(ctx *gin.Context) {
	var input struct {
		Data []primitive.ObjectID `json:"data" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Approve(input.Data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAll)
	publicRg.POST("/create", this.create)

	adminRg.PATCH("/update", this.update)
	adminRg.PATCH("/approve", this.approve)
	adminRg.DELETE("/delete", this.delete)
}
