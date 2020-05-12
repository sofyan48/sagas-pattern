package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetToken params
// @context: *gin.Context
// return gin.HandlerFunc
func (m *DefaultMiddleware) GetToken(context *gin.Context) string {
	token := context.Request.Header["Authorization"]
	if len(token) < 1 {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Token Not Compatible",
		})
		context.Abort()
		return "event"
	}
	return token[0]
}

// AuthToken ...
func (m *DefaultMiddleware) AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := m.GetToken(ctx)
		fmt.Println(tokenString)
	}
}
