package utils

import (
	"aie/internal/model"
	"errors"
	"os"

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
