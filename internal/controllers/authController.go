package controllers

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"strings"
	"urfu-radio-journal/internal/models"
	"urfu-radio-journal/pkg/services"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	auth *services.AuthService
}

func NewAuthController(frontendAdress string) *AuthController {
	frontendDomain := strings.Split(frontendAdress, ":")[1][2:]
	return &AuthController{auth: services.NewAuthService(frontendDomain)}
}

func (this *AuthController) login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := this.auth.Login(admin, session); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *AuthController) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if err := this.auth.Logout(session); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (this *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.Use(sessions.Sessions("admin", this.auth.GetStore()))

	rg.POST("/login", this.login)
	rg.GET("/logout", this.logout)
}

func (this *AuthController) AuthMiddleware() gin.HandlerFunc {
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

		session := sessions.Default(ctx)
		if admin := session.Get("admin"); admin == nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func (this *AuthController) SessionsHandler() gin.HandlerFunc {
	return sessions.Sessions("admin", this.auth.GetStore())
}
