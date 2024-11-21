package server

import (
	"aie/internal/api/authentication"
	"aie/internal/api/discussion"
	"aie/internal/api/group"
	"aie/internal/api/onboarding"
	"aie/internal/api/profile"
	"aie/internal/api/sockets"
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
		discussion.DiscussionAPIs(v1)
		group.GroupAPIs(v1)
		sockets.SocketAPIs(v1)
	}
}
