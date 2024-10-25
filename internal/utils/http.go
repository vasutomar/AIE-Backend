package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PrepareResponse - Prepare response for the API
func SetResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, gin.H{"message": message, "statusCode": statusCode, "data": data, "error": nil})
}

func SetError(c *gin.Context, err error) {
	c.JSON(http.StatusInternalServerError, gin.H{"message": "We hit a snag!", "statusCode": http.StatusInternalServerError, "data": nil, "error": err.Error()})
}
