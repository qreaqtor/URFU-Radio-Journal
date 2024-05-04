package councilhand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.CouncilMemberCreate) error
	GetAll() ([]*models.CouncilMemberRead, error)
	Get(string) (*models.CouncilMemberRead, error)
	Update(string, *models.CouncilMemberUpdate) error
	Delete(string) error
}

type CouncilHandler struct {
	members service
}

func NewCouncilHandler(council service) *CouncilHandler {
	return &CouncilHandler{
		members: council,
	}
}

func (c *CouncilHandler) Create(ctx *gin.Context) {
	member := &models.CouncilMemberCreate{}
	if err := ctx.ShouldBindJSON(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := c.members.Create(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (c *CouncilHandler) Update(ctx *gin.Context) {
	member := &models.CouncilMemberUpdate{}
	if err := ctx.ShouldBindJSON(member); err != nil {
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

func (c *CouncilHandler) GetAll(ctx *gin.Context) {
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

func (c *CouncilHandler) GetMemberById(ctx *gin.Context) {
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

func (c *CouncilHandler) Delete(ctx *gin.Context) {
	memberId := ctx.Param("id")
	// memberId, err := primitive.ObjectIDFromHex(memberIdStr)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// imagePathId, err := c.members.GetImagePathId(memberId)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	// filter := bson.M{"_id": imagePathId}
	// if err := c.deleteFiles(filter); err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	// 	return
	// }
	if err := c.members.Delete(memberId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

// func (c *CouncilHandler) RegisterRoutes(publicRg, adminRg *gin.RouterGroup) {
// 	publicRg.GET("/get/all", c.getAll)
// 	publicRg.GET("/get/:memberId", c.getMemberById)

// 	adminRg.POST("/create", c.create)
// 	adminRg.PUT("/update/:id", c.update)
// 	adminRg.DELETE("/delete/:id", c.delete)
// }
