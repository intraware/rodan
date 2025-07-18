package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/intraware/rodan/utils"
	"github.com/intraware/rodan/utils/values"
)

var AuthRequired gin.HandlerFunc = func(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		ctx.Abort()
		return
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token required"})
		ctx.Abort()
		return
	}
	claims, err := utils.ValidateJWT(tokenString, values.GetConfig().Security.JWTSecret)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}
	ctx.Set("user_id", claims.UserID)
	ctx.Set("username", claims.Username)
	ctx.Next()
}
