package server

import (
	"aie/internal/api/authentication"
	"aie/internal/api/onboarding"
	"aie/internal/api/profile"
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine) {
	router.Use(gin.Logger())
	router.Use(middlewares.CORSMiddleware())
	v1 := router.Group("/api/v1/")
	{
		authentication.AuthenticationAPIs(v1)
		onboarding.OnboardingAPIs(v1)
		profile.ProfileAPIs(v1)
	}
}
