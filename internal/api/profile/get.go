package profile

import (
	"aie/internal/model"
	"aie/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Note: Current with own's token, user will be able to get his/her + any other user's profile
func GetProfile(c *gin.Context) {
	userId := utils.GetUserId(c)
	profile, err := model.GetProfile(userId)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Profile fetched", profile)
}

func GetFriendsForUser(c *gin.Context) {
	userId := utils.GetUserId(c)
	friends, err := model.GetFriends(userId)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Friends fetched", friends)

}

func GetProfilesBasedOnExam(c *gin.Context) {
	exam := c.Param("exam")
	itemsCount := c.DefaultQuery("items", "20")
	page := c.DefaultQuery("page", "1")

	itemsCountInt, err := strconv.ParseInt(itemsCount, 10, 64)
	if err != nil {
		utils.SetError(c, err)
		return
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		utils.SetError(c, err)
		return
	}

	profiles, _ := model.GetAllProfilesForAnExam(exam, itemsCountInt, pageInt)
	utils.SetResponse(c, 200, "Profiles fetched", profiles)
}

func ProfileSearch(c *gin.Context) {
	exam := c.Param("exam")
	name := c.Param("name")
	profiles, _ := model.GetUsersByName(exam, name)
	utils.SetResponse(c, 200, "Profiles fetched", profiles)
}
