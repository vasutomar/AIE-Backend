package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDiscussionsByExam(c *gin.Context) {
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

	discussions, err := model.GetDiscussionsByExam(exam, itemsCountInt, pageInt)
	utils.SetResponse(c, 200, "Discussions fetched", discussions)
}
