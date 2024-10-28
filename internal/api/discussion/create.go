package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDiscussion(c *gin.Context) {
	discussion := model.Discussion{}

	if err := c.ShouldBindJSON(&discussion); err != nil {
		utils.SetError(c, err)
		return
	}

	if err := discussion.Create(); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Discussion created up successfully", nil)
}
