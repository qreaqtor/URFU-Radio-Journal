package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func PanicMiddleware() gin.HandlerFunc {
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
		ctx.Next()
	}
}
