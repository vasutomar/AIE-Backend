package todo

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

// Note: Current with own's token, user will be able to get his/her + any other user's profile
func GetTodo(c *gin.Context) {
	userId := utils.GetUserId(c)
	profile, err := model.GetTodoData(userId)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 201, "Todo fetched", profile)
}
