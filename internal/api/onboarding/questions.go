package onboarding

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Questions(c *gin.Context) {
	questions, err := model.GetQuestions()
	if err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Questions fetched", questions)
}
