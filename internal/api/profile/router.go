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
		r.PATCH("", middlewares.AuthMiddleware(), UpdateProfile)
	}
}
