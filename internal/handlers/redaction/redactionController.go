package redactionhand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Create(*models.RedactionMemberCreate) error
	GetAll() ([]*models.RedactionMemberRead, error)
	Get(string) (*models.RedactionMemberRead, error)
	Update(string, *models.RedactionMemberUpdate) error
	Delete(string) error
}

type RedactionHandler struct {
	members service
}

func NewRedactionHandler(council service) *RedactionHandler {
	return &RedactionHandler{
		members: council,
	}
}

func (r *RedactionHandler) Create(ctx *gin.Context) {
	member := &models.RedactionMemberCreate{}
	if err := ctx.ShouldBindJSON(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := r.members.Create(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (r *RedactionHandler) Update(ctx *gin.Context) {
	member := &models.RedactionMemberUpdate{}
	if err := ctx.ShouldBindJSON(member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	memberId := ctx.Param("id")
	if err := r.members.Update(memberId, member); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (r *RedactionHandler) GetAll(ctx *gin.Context) {
	data, err := r.members.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"data":    data,
		"message": "success",
	})
}

func (r *RedactionHandler) GetMemberById(ctx *gin.Context) {
	memberId := ctx.Param("memberId")
	member, err := r.members.Get(memberId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"member":  member,
		"message": "success",
	})
}

func (r *RedactionHandler) Delete(ctx *gin.Context) {
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
	if err := r.members.Delete(memberId); err != nil {
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
