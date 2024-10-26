package discussion

import (
	"aie/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func DiscussionAPIs(router *gin.RouterGroup) {
	r := router.Group("/discussion/")
	{
		r.GET("health", Health)
		r.GET(":exam", middlewares.AuthMiddleware(), GetDiscussionsByExam)
	}
}
