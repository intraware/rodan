package auth

import (
	"time"

	"github.com/AnimeKaizoku/cacher"
	"github.com/gin-gonic/gin"
	"github.com/intraware/rodan/api/shared"
	"github.com/intraware/rodan/models"
	"github.com/intraware/rodan/utils/values"
)

func LoadAuth(r *gin.RouterGroup) {
	shared.ResetPasswordCache = cacher.NewCacher[string, models.User](&cacher.NewCacherOpts{
		TimeToLive:    time.Duration(values.GetConfig().App.TokenExpiry) * time.Minute,
		CleanInterval: time.Hour * 2,
		CleanerMode:   cacher.CleaningCentral,
	})

	authRouter := r.Group("/auth")

	authRouter.POST("/signup", signUp)
	authRouter.POST("/login", login)
	authRouter.POST("/forgot-password", forgotPassword)
	authRouter.POST("/reset-password/:token", resetPassword)
}
