package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CheckApproved(ctx *gin.Context) {
	approved, err := strconv.ParseBool(ctx.Query("onlyApproved"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Set("skip", approved)
	ctx.Next()
}
