package todo

import (
	"aie/internal/model"
	"aie/internal/utils"

	"github.com/gin-gonic/gin"
)

// Note: Current with own's token, user will be able to update his/her + any other user's profile
func UpdateTodo(c *gin.Context) {
	todoId := c.Param("id")
	todo := &model.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.SetError(c, err)
		return
	}

	err := model.UpdateTodo(todo, todoId)
	if err != nil {
		utils.SetError(c, err)
		return
	}
	utils.SetResponse(c, 200, "Todo updated", nil)
}
