package controllers

import (
	"net/http"
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
	onlyApproved := true
	onlyApprovedStr := ctx.Query("onlyApproved")
	if onlyApprovedStr == "false" {
		onlyApproved = false
	}
	articleId := ctx.Query("articleId")
	comments, err := this.comments.GetAll(onlyApproved, articleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": comments})
}

func (this *CommentsController) update(ctx *gin.Context) {
	var comment models.CommentUpdate
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Update(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) delete(ctx *gin.Context) {
	commentIdStr := ctx.Param("id")
	commentId, err := primitive.ObjectIDFromHex(commentIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Delete(commentId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) approve(ctx *gin.Context) {
	var commentApprove models.CommentApprove
	if err := ctx.ShouldBindJSON(&commentApprove); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.comments.Approve(commentApprove); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *CommentsController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAll)
	publicRg.POST("/create", this.create)

	adminRg.PATCH("/update", this.update)
	adminRg.PATCH("/approve", this.approve)
	adminRg.DELETE("/delete/:id", this.delete)
}

func (this *CommentsController) GetDeleteHandler() func(filter primitive.M) error {
	return this.comments.DeleteManyHandler
}
