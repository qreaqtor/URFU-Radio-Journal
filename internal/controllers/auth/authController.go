package auth

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services/auth"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth *auth.AuthService
}

func NewAuthController() *AuthController {
	return &AuthController{auth: auth.NewAuthService()}
}

func (a *AuthController) login(ctx *gin.Context) {
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

// func (this *AuthController) logout(ctx *gin.Context) {
// 	session := sessions.Default(ctx)
// 	if err := this.auth.Logout(session); err != nil {
// 		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
// }

func (a *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.POST("/login", a.login)
	// rg.GET("/logout", this.logout)
}

func (a *AuthController) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Panic: %v\n", r)
				debug.PrintStack()
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
				ctx.Abort()
				return
			}
		}()

		token := extractToken(ctx)
		err := a.auth.ValidateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("can't extract token: %s", err.Error())})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func extractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}
