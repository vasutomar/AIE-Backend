package utils

import (
	"aie/internal/model"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyUserJWT(tokenString string) (*model.UserToken, error) {
	secretKey := []byte(os.Getenv("SALT"))
	// Parse the token and validate it using the secret key
	token, err := jwt.ParseWithClaims(tokenString, &model.UserToken{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*model.UserToken); ok && token.Valid {
		if claims.Issuer != "aie-backend-service" {
			return nil, errors.New("Invalid token issuer")
		}
		return claims, nil
	} else {
		return nil, err
	}
}

func GetUserId(c *gin.Context) string {

	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
		c.Abort()
		return ""
	}

	// The token is typically prefixed by "Bearer"
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		c.Abort()
		return ""
	}

	claims, _ := VerifyUserJWT(tokenString)
	return claims.UserId
}
