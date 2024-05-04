package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(validateToken func(string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := extractToken(ctx)
		err := validateToken(token)
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
