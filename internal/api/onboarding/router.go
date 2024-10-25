package onboarding

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func OnboardingAPIs(router *gin.RouterGroup) {
	r := router.Group("/onboarding/")
	{
		r.GET("health", Health)
		r.GET("questions", middlewares.AuthMiddleware(), Questions)
	}
}
