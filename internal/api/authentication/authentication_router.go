package authentication

import (
	"github.com/gin-gonic/gin"
)

func AuthenticationAPIs(router *gin.RouterGroup) {
	authRoute := router.Group("/authentication/")
	{
		authRoute.GET("health", Health)
		authRoute.POST("signup", Signup)
		authRoute.POST("signin", Signin)
		authRoute.POST("verify", Verify)
	}
}
