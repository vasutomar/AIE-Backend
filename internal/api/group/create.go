package group

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {
	createGroupRequestData := model.CreateGroupRequest{}

	if err := c.ShouldBindJSON(&createGroupRequestData); err != nil {
		utils.SetError(c, err)
		return
	}
	userId := utils.GetUserId(c)
	if err := createGroupRequestData.Create(userId); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Group created up successfully", nil)
}
