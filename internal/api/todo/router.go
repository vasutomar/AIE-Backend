package todo

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func TodoAPIs(router *gin.RouterGroup) {
	r := router.Group("/todo/")
	{
		r.GET("health", Health)
		r.GET("", middlewares.AuthMiddleware(), GetTodo)
		r.POST("", middlewares.AuthMiddleware(), CreateTodo)
		r.PATCH(":id", middlewares.AuthMiddleware(), UpdateTodo)
	}
}
