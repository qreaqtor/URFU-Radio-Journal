package comments

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services/comments"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentsController struct {
	comments *comments.CommentsService
}

func NewCommentsController() *CommentsController {
	return &CommentsController{comments: comments.NewCommentsService()}
}

func (c *CommentsController) create(ctx *gin.Context) {
	var comment models.CommentCreate
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.comments.Create(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CommentsController) getAll(ctx *gin.Context) {
	onlyApproved := true
	onlyApprovedStr := ctx.Query("onlyApproved")
	if onlyApprovedStr == "false" {
		onlyApproved = false
	}
	articleId := ctx.Query("articleId")
	comments, err := c.comments.GetAll(onlyApproved, articleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": comments})
}

func (c *CommentsController) update(ctx *gin.Context) {
	var comment models.CommentUpdate
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.comments.Update(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CommentsController) delete(ctx *gin.Context) {
	commentIdStr := ctx.Param("id")
	commentId, err := primitive.ObjectIDFromHex(commentIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.comments.Delete(commentId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CommentsController) approve(ctx *gin.Context) {
	var commentApprove models.CommentApprove
	if err := ctx.ShouldBindJSON(&commentApprove); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.comments.Approve(commentApprove); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CommentsController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", c.getAll)
	publicRg.POST("/create", c.create)

	adminRg.PATCH("/update", c.update)
	adminRg.PATCH("/approve", c.approve)
	adminRg.DELETE("/delete/:id", c.delete)
}

func (c *CommentsController) GetDeleteHandler() func(filter primitive.M) error {
	return c.comments.DeleteManyHandler
}
