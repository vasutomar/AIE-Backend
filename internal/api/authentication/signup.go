package authentication

import (
	"aie/internal/model"
	"aie/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Signup(c *gin.Context) {
	userData := model.UserSignupRequest{}
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userProfile := model.Profile{
		Username: userData.Username,
		Phone:    userData.Phone,
		Email:    userData.Email,
		Exams:    []string{},
		Salt:     "",
	}

	user := model.User{
		Username:  userData.Username,
		Password:  userData.Password,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		UserId:    userData.UserId,
	}

	jwt, err := user.Create()
	if err != nil {
		log.Err(err).Msg("Error creating user")
		utils.SetError(c, err)
		return
	}
	profileCreateError := model.CreateProfile(userProfile)
	if profileCreateError != nil {
		log.Err(err).Msg("Error creating user profile")
		utils.SetError(c, err)
		return
	}

	utils.SetResponse(c, http.StatusOK, "User signed up successfully", jwt)
}
