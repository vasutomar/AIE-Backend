package authentication

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Signup(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := user.Create()
	if err != nil {
		log.Err(err).Msg("Error creating user")
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "User signed up successfully", jwt)
}
