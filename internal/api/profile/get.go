package profile

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

// Note: Current with own's token, user will be able to get his/her + any other user's profile
func GetProfile(c *gin.Context) {
	username := c.Param("username")
	profile, err := model.GetProfile(username)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Profile fetched", profile)
}
