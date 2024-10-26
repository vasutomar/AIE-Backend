package discussion

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

func GetDiscussionsByExam(c *gin.Context) {
	exam := c.Param("exam")
	profile, err := model.GetDiscussions(exam)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Discussions fetched", profile)
}
