package util

import (
	"crypto/sha256"
	"fmt"

	"github.com/Fadhelbulloh/api-with-redis/model/response"
	"github.com/gin-gonic/gin"
)

// Hashing return key for redis from hashed byte parameter and key API
func Hashing(stringParam []byte, key string) string {
	hash := sha256.New()
	hash.Write(stringParam)
	hashQuery := fmt.Sprintf("%s-%x", key, hash.Sum(nil))

	return hashQuery
}

// ErrHandler handling error for gin controller, return true if there's error
func ErrHandler(code int, c *gin.Context, err error) bool {
	if err != nil {
		c.JSON(code, gin.H{"status": false, "message": err.Error(), "data": nil})
		return true
	}
	return false
}

// ErrHandler handling response error for gin controlleri, return true if there's error
func ErrorHandleResponse(c *gin.Context, response response.Response) bool {
	if !response.Status {
		c.JSON(response.StatusCode, response)
		return true
	}
	return false
}
