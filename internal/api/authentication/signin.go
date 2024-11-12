package authentication

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Signin(c *gin.Context) {
	user := model.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// This user will only have username and password fields set.
	responseData, err := user.Signin()
	if err != nil {
		log.Err(err).Msg("Error signin for user")
		utils.SetResponse(c, 204, "Invalid username or password", nil)
		return
	}

	utils.SetResponse(c, http.StatusOK, "User signed in successfully", responseData)
}
