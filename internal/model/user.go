package model

import (
	"aie/internal/providers"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserId    string `json:"userid"`
}

type UserToken struct {
	UserId    string
	Username  string
	FirstName string
	LastName  string
	jwt.RegisteredClaims
}

type UserSigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignupRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	UserId    string `json:"userid"`
}

type JWTTokenData struct {
	Username  string `json:"username"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	UserId    string `json:"userid"`
}

type UserTokenRequest struct {
	Token string `json:"token"`
}

// TODO: Add timeouts to the context
func (user *User) Signin() (string, error) {
	log.Debug().Msgf("user Signin started: %v", user)
	user.hashPassword()
	filter := bson.D{{Key: "username", Value: user.Username}, {Key: "password", Value: user.Password}}

	var result bson.M
	err := providers.DB.Collection("USERS").FindOne(context.Background(), filter).Decode(&result)
	exists := true
	if err != nil {
		log.Debug().Msgf("Error finding user: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			exists = false
		} else {
			return "", err
		}
	}

	if exists {
		tokenData := JWTTokenData{
			FirstName: result["firstname"].(string),
			LastName:  result["lastname"].(string),
			UserId:    result["userid"].(string),
			Username:  result["username"].(string),
		}
		jwt, err := tokenData.generateJWT()
		if err != nil {
			return "", err
		}

		log.Debug().Msgf("User signed in successfully: %v, jwt=%s", user, jwt)
		return jwt, nil
	}
	return "", errors.New("User does not exist")
}

// TODO: Add timeouts to the context
// Note: The field names stored in DB will be in small case always.
func (user *User) Create() (string, error) {
	log.Debug().Msgf("user Create started: %v", user)
	// Check if the user already exists
	filter := bson.D{{Key: "username", Value: user.Username}}

	var result bson.M
	err := providers.DB.Collection("USERS").FindOne(context.Background(), filter).Decode(&result)
	exists := true
	if err != nil {
		log.Debug().Msgf("Error finding user: %v", err)
		if errors.Is(err, mongo.ErrNoDocuments) {
			exists = false
		} else {
			return "", err
		}
	}

	if !exists {
		// Create user
		user.UserId = uuid.New().String()
		user.hashPassword()
		log.Debug().Msgf("Creating user: %v", user)
		_, err = providers.DB.Collection("USERS").InsertOne(context.Background(), user)
		if err != nil {
			return "", err
		}

		tokenData := JWTTokenData{
			Username:  user.Username,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			UserId:    user.UserId,
		}

		jwt, err := tokenData.generateJWT()
		if err != nil {
			return "", err
		}

		log.Debug().Msgf("User created successfully: %v, jwt=%s", user, jwt)
		return jwt, nil
	}

	return "", errors.New(fmt.Sprintf("User already exists: %v", user.Username))
}

func (data *JWTTokenData) generateJWT() (string, error) {
	// TODO: Add test for JWT expiry
	claims := UserToken{
		data.Username,
		data.FirstName,
		data.LastName,
		data.UserId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "aie-backend-service",
		},
	}

	// Create a new JWT token with the claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key (salt)
	tokenString, err := token.SignedString([]byte(os.Getenv("SALT")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (user *User) hashPassword() {
	// Define the secret key
	secretKey := os.Getenv("PASSALGO")

	// Create a new HMAC by defining the hash type and the secret key
	h := hmac.New(sha256.New, []byte(secretKey))

	// Write the password data to the HMAC object
	h.Write([]byte(user.Password))

	// Compute the HMAC and return the result as a hexadecimal string
	user.Password = hex.EncodeToString(h.Sum(nil))
}
