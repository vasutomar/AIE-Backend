package todo

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	createTodoRequestData := model.CreateTodoRequest{}

	if err := c.ShouldBindJSON(&createTodoRequestData); err != nil {
		utils.SetError(c, err)
		return
	}
	userId := utils.GetUserId(c)
	if err := createTodoRequestData.Create(userId); err != nil {
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Todo created successfully", nil)
}
