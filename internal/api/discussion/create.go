package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDiscussion(c *gin.Context) {
	createDiscussionRequestData := model.CreateDiscusstionRequest{}

	if err := c.ShouldBindJSON(&createDiscussionRequestData); err != nil {
		utils.SetError(c, err)
		return
	}
	userId := utils.GetUserId(c)
	if err := createDiscussionRequestData.Create(userId); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Discussion created up successfully", nil)
}
