package group

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func GroupAPIs(router *gin.RouterGroup) {
	r := router.Group("/group/")
	{
		r.GET("health", Health)
		r.GET(":id", middlewares.AuthMiddleware(), GetGroup)
		r.GET("", middlewares.AuthMiddleware(), GetGroupsForUser)
		r.POST("", middlewares.AuthMiddleware(), CreateGroup)
	}
}
