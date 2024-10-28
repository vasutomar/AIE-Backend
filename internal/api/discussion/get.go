package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDiscussions(c *gin.Context) {
	exam := c.Param("exam")
	count := c.Query("count")
	page := c.Query("page")

	countInt, err := strconv.ParseInt(count, 10, 32)
	if err != nil {
		utils.SetError(c, err)
		return
	}

	pageInt, err := strconv.ParseInt(page, 10, 32)
	if err != nil {
		utils.SetError(c, err)
		return
	}

	discussions, err := model.GetDiscussions(exam, countInt, pageInt)
	utils.SetResponse(c, 200, "Discussions fetched", discussions)
}
