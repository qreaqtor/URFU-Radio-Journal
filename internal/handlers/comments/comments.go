package commentshand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.CommentCreate) error
	GetAll(*models.CommentQuery) ([]*models.CommentRead, int, error)
	Update(*models.CommentUpdate) error
	Delete(string) error
	Approve(*models.CommentApprove) error
}

type CommentsHandler struct {
	comments service
}

func NewCommentsHandler(comments service) *CommentsHandler {
	return &CommentsHandler{
		comments: comments,
	}
}

func (c *CommentsHandler) Create(ctx *gin.Context) {
	comment := &models.CommentCreate{}
	if err := ctx.ShouldBindJSON(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.comments.Create(comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CommentsHandler) GetAll(ctx *gin.Context) {
	args := &models.CommentQuery{}
	err := ctx.ShouldBindQuery(args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	comments, count, err := c.comments.GetAll(args)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":      comments,
		"all_count": count,
	})
}

func (c *CommentsHandler) Update(ctx *gin.Context) {
	comment := &models.CommentUpdate{}
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

func (c *CommentsHandler) Delete(ctx *gin.Context) {
	commentId := ctx.Param("id")
	if err := c.comments.Delete(commentId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CommentsHandler) Approve(ctx *gin.Context) {
	commentApprove := &models.CommentApprove{}
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
