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
		r.POST("", middlewares.AuthMiddleware(), CreateDiscussion)
		r.PATCH(":id", middlewares.AuthMiddleware(), UpdateDiscussion)
		r.PATCH("/comment/:id", middlewares.AuthMiddleware(), CommentOnDiscussion)
		r.PATCH("/like/:id", middlewares.AuthMiddleware(), ToggleDiscussionLike)
		r.PATCH("/bookmark/:id", middlewares.AuthMiddleware(), ToggleDiscussionBookmark)
	}
}
