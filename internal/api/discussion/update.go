package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

func UpdateDiscussion(c *gin.Context) {
	discussion := &model.Discussion{}
	if err := c.ShouldBindJSON(&discussion); err != nil {
		utils.SetError(c, err)
		return
	}

	id := c.Param("id")
	err := model.UpdateDiscussion(id, discussion)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 200, "Discussion updated", nil)
}
