package onboarding

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func OnboardingAPIs(router *gin.RouterGroup) {
	authRoute := router.Group("/onboarding/")
	{
		authRoute.GET("health", Health)
		authRoute.GET("questions", middlewares.AuthMiddleware(), Questions)
	}
}
