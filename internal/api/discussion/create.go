package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateDiscussion(c *gin.Context) {
	discussion := &model.Discussion{}
	if err := c.ShouldBindJSON(&discussion); err != nil {
		utils.SetError(c, err)
		return
	}
	err := model.CreateDiscussion(*discussion)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Discussion created", nil)
}
