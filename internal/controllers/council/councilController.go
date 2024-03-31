package council

import (
	"net/http"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services/council"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CouncilController struct {
	members     *council.CouncilService
	deleteFiles func(filter primitive.M) error
}

func NewCouncilController(deleteFilesHandler func(filter primitive.M) error) *CouncilController {
	return &CouncilController{
		members:     council.NewCouncilService(),
		deleteFiles: deleteFilesHandler,
	}
}

func (c *CouncilController) create(ctx *gin.Context) {
	var member models.CouncilMemberCreate
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.members.Create(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CouncilController) update(ctx *gin.Context) {
	var member models.CouncilMemberUpdate
	if err := ctx.ShouldBindJSON(&member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	memberId := ctx.Param("id")
	if err := c.members.Update(memberId, member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CouncilController) getAll(ctx *gin.Context) {
	data, err := c.members.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

func (c *CouncilController) getMemberById(ctx *gin.Context) {
	memberId := ctx.Param("memberId")
	member, err := c.members.Get(memberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"member":  member,
		"message": "success",
	})
}

func (c *CouncilController) delete(ctx *gin.Context) {
	memberIdStr := ctx.Param("id")
	memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	imagePathId, err := c.members.GetImagePathId(memberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	filter := bson.M{"_id": imagePathId}
	if err := c.deleteFiles(filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.members.Delete(memberId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CouncilController) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
	publicRg.GET("/get/all", c.getAll)
	publicRg.GET("/get/:memberId", c.getMemberById)

	adminRg.POST("/create", c.create)
	adminRg.PUT("/update/:id", c.update)
	adminRg.DELETE("/delete/:id", c.delete)
}