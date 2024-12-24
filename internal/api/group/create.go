package group

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGroup(c *gin.Context) {
	createGroupRequestData := model.CreateGroupRequest{}
	responseData := make(map[string]string)

	if err := c.ShouldBindJSON(&createGroupRequestData); err != nil {
		utils.SetError(c, err)
		return
	}
	userId := utils.GetUserId(c)
	creatorProfile, _ := model.GetCondensedProfile(userId)
	createGroupRequestData.Members = append(createGroupRequestData.Members, creatorProfile)
	if group_id, err := createGroupRequestData.Create(userId); err != nil {
		utils.SetError(c, err)
		return
	} else {
		profile := model.Profile{
			Groups: []string{group_id},
		}
		for _, member := range createGroupRequestData.Members {
			responseData["group_id"] = group_id
			_ = model.UpdateProfile(member.UserId, &profile)
		}
	}

	utils.SetResponse(c, http.StatusOK, "Group created up successfully", responseData)
}
