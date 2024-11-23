package profile

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileAPIs(router *gin.RouterGroup) {
	r := router.Group("/profile/")
	{
		r.GET("health", Health)
		r.GET("", middlewares.AuthMiddleware(), GetProfile)
		r.GET("friends", middlewares.AuthMiddleware(), GetFriendsForUser)
		r.PATCH("", middlewares.AuthMiddleware(), UpdateProfile)
	}
}
