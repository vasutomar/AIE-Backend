package authentication

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Verify(c *gin.Context) {
	userToken := model.UserTokenRequest{}

	if err := c.ShouldBindJSON(&userToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := utils.VerifyUserJWT(userToken.Token)
	if err != nil {
		log.Err(err).Msg("Error verifying token for user")
		utils.SetResponse(c, 401, "Invalid session", nil)
		return
	}

	utils.SetResponse(c, http.StatusOK, "Verified session", nil)
}
