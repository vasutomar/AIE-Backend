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
		r.GET("all/:exam", middlewares.AuthMiddleware(), GetProfilesBasedOnExam)
		r.GET("user/:name/:exam", middlewares.AuthMiddleware(), GetProfilesBasedOnExam)
		r.PATCH("", middlewares.AuthMiddleware(), UpdateProfile)
	}
}
