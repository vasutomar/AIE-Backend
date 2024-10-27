package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDiscussionsByExam(c *gin.Context) {
	exam := c.Param("exam")
	items, _ := strconv.ParseInt(c.DefaultQuery("items", "1"), 10, 64)
	page, _ := strconv.ParseInt(c.DefaultQuery("page", "1"), 10, 64)
	profile, err := model.GetDiscussions(exam, items, page)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Discussions fetched", profile)
}
