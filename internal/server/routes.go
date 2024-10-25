package server

import (
	"aie/internal/api/authentication"
	"aie/internal/api/onboarding"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	v1 := router.Group("/api/v1/")
	{
		authentication.AuthenticationAPIs(v1)
		onboarding.OnboardingAPIs(v1)
	}
}
