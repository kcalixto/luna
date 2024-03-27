package responseUtils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, message ...string) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": message})
}

func Success(ctx *gin.Context, message ...string) {
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}

func Created(ctx *gin.Context, message ...string) {
	ctx.JSON(http.StatusCreated, gin.H{"message": message})
}

func NoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func InternalServerError(ctx *gin.Context, message ...string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": message})
}

func NotFound(ctx *gin.Context, message ...string) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": message})
}

func Unauthorized(ctx *gin.Context, message ...string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{"error": message})
}
