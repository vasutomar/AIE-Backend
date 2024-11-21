package sockets

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SocketAPIs(router *gin.RouterGroup) {
	r := router.Group("/socket/")
	{
		r.GET("health", Health)
		r.GET("connect/:groupId", middlewares.AuthMiddleware(), Connect)
	}
}
