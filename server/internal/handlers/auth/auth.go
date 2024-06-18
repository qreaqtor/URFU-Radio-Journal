package authand

import (
	"net/http"
	"urfu-radio-journal/internal/models"

	"github.com/gin-gonic/gin"
)

type service interface {
	Login(models.Admin) (string, error)
}

type AuthHandler struct {
	auth service
}

func NewAuthHandler(auth service) *AuthHandler {
	return &AuthHandler{
		auth: auth,
	}
}

func (a *AuthHandler) Login(ctx *gin.Context) {
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	token, err := a.auth.Login(admin)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
		"Token":   token,
	})
}
