package profile

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

// Note: Current with own's token, user will be able to update his/her + any other user's profile
func UpdateProfile(c *gin.Context) {
	profile := &model.Profile{}
	if err := c.ShouldBindJSON(&profile); err != nil {
		utils.SetError(c, err)
		return
	}

	userId := utils.GetUserId(c)
	err := model.UpdateProfile(userId, profile)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 200, "Profile updated", nil)
}
