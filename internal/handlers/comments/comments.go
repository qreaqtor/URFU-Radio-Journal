package commentshand

import (
	"net/http"
	"strconv"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.CommentCreate) error
	GetAll(bool, string) ([]*models.CommentRead, error)
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
	onlyApproved, err := strconv.ParseBool(ctx.Query("onlyApproved"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	articleId := ctx.Query("articleId")
	comments, err := c.comments.GetAll(onlyApproved, articleId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": comments})
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
	// commentId, err := primitive.ObjectIDFromHex(commentIdStr)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
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

// func (c *CommentsHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
// 	publicRg.GET("/get/all", c.getAll)
// 	publicRg.POST("/create", c.create)

// 	adminRg.PATCH("/update", c.update)
// 	adminRg.PATCH("/approve", c.approve)
// 	adminRg.DELETE("/delete/:id", c.delete)
// }

// func (c *CommentsHandler) GetDeleteHandler() func(filter primitive.M) error {
// 	return c.comments.DeleteManyHandler
// }
