package controllers

import (
	"urfu-radio-journal/pkg/services"

	"github.com/gin-gonic/gin"
)

type CommentsController struct {
	comments *services.CommentsService
}

func NewCommentsController() *CommentsController {
	return &CommentsController{comments: services.NewCommentsService()}
}

func (this *CommentsController) create(ctx *gin.Context) {

}

func (this *CommentsController) getAll(ctx *gin.Context) {

}

func (this *CommentsController) update(ctx *gin.Context) {

}

func (this *CommentsController) delete(ctx *gin.Context) {

}

func (this *CommentsController) RegisterRoutes(publicRg *gin.RouterGroup, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", this.getAll)
	publicRg.POST("/create", this.create)

	adminRg.PATCH("/update", this.update)
	adminRg.DELETE("/delete", this.delete)
}
