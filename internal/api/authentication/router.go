package authentication

import (
	"github.com/gin-gonic/gin"
)

func AuthenticationAPIs(router *gin.RouterGroup) {
	r := router.Group("/authentication/")
	{
		r.GET("health", Health)
		r.POST("signup", Signup)
		r.POST("signin", Signin)
		r.POST("verify", Verify)
	}
}
