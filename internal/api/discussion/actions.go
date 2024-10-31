package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CommentOnDiscussion(c *gin.Context) {
	commentData := model.CommentRequest{}
	id := c.Param("id")
	if err := c.ShouldBindJSON(&commentData); err != nil {
		utils.SetError(c, err)
		return
	}
	username := utils.GetUsername(c)
	comment := model.Comment{
		Username: username,
		Comment:  commentData.Comment,
	}
	if err := model.AddComment(id, comment, username); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Comment added successfully", nil)
}

func ToggleDiscussionLike(c *gin.Context) {
	commentData := model.CommentRequest{}
	discussion_id := c.Param("id")
	if err := c.ShouldBindJSON(&commentData); err != nil {
		utils.SetError(c, err)
		return
	}
	user_id := utils.GetUserId(c)
	if err := model.ToggleLike(user_id, discussion_id); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Like toggled successfully", nil)
}

func ToggleDiscussionBookmark(c *gin.Context) {
	commentData := model.CommentRequest{}
	discussion_id := c.Param("id")
	if err := c.ShouldBindJSON(&commentData); err != nil {
		utils.SetError(c, err)
		return
	}
	user_id := utils.GetUserId(c)
	if err := model.ToggleBookmark(user_id, discussion_id); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Bookmark toggled successfully", nil)
}
