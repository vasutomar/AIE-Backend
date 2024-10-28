package middlewares

import (
	"aie/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Middleware to verify JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// The token is typically prefixed by "Bearer"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		claims, err := utils.VerifyUserJWT(tokenString)

		// Check if token is valid
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// You can pass the claims down to the handler using the context
		c.Set("username", claims.Username)
		c.Set("userId", claims.UserId)
		c.Set("firstName", claims.FirstName)
		c.Set("lastName", claims.LastName)

		// Proceed with the request
		c.Next()
	}
}
