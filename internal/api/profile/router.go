package profile

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileAPIs(router *gin.RouterGroup) {
	r := router.Group("/profile/")
	{
		r.GET("health", Health)
		r.GET(":username", middlewares.AuthMiddleware(), GetProfile)
		r.PATCH(":username", middlewares.AuthMiddleware(), UpdateProfile)
	}
}
