package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	errEmptyToken = errors.New("empty token")
)

func AuthMiddleware(validateToken func(string) error) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := extractToken(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("no token: %s", err.Error())})
			ctx.Abort()
			return
		}
		err = validateToken(token)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("bad token: %s", err.Error())})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

func extractToken(ctx *gin.Context) (string, error) {
	bearerToken := ctx.Request.Header.Get("Authorization")
	token := strings.TrimPrefix(bearerToken, "Bearer: ")
	if token == "" {
		return "", errEmptyToken
	}
	return token, nil
}
