package controllers

import (
	"fmt"
	"net/http"
	"runtime/debug"
	"urfu-radio-journal/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	store cookie.Store
}

func NewAuthController() *AuthController {
	store := cookie.NewStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge:   3600, // seconds
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // only for HTTPS
	})

	return &AuthController{store: store}
}

func checkAdmin(admin models.Admin) bool {
	return admin.Username == "admin" && admin.Password == "admin"
}

func (this *AuthController) login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	var admin models.Admin
	if err := ctx.ShouldBindJSON(&admin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if checkAdmin(admin) {
		session.Set("admin", admin.Username)
		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "can't save session"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
}

func (this *AuthController) logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if admin := session.Get("admin"); admin != nil {
		session.Delete("admin")
		if err := session.Save(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "can't save session"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
		return
	}
	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
}

func (this *AuthController) RegisterRoutes(rg *gin.RouterGroup) {
	rg.Use(sessions.Sessions("admin", this.store))

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
	return sessions.Sessions("admin", this.store)
}
