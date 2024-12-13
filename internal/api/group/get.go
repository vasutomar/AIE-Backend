package group

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGroupsForUser(c *gin.Context) {
	userId := utils.GetUserId(c)
	profile, err := model.GetProfile(userId)
	if err != nil {
		utils.SetError(c, err)
		return
	}

	groups := profile.Groups
	fetchedGroupsForUser, _ := model.GetGroups(groups)
	utils.SetResponse(c, http.StatusOK, "Groups fetched successfully", fetchedGroupsForUser)
}

func GetGroup(c *gin.Context) {
	groupId := c.Param("id")
	fetchedGroup, _ := model.GetGroup(groupId)
	utils.SetResponse(c, http.StatusOK, "Group fetched successfully", fetchedGroup)
}
