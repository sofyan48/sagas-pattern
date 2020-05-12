package middleware

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_user/src/utils/requester"
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
func (m *DefaultMiddleware) AuthToken(target string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := m.GetToken(ctx)
		uriOauth := os.Getenv("OAUTH_URL") + "/v1/authorization/client"
		if target != "" {
			uriOauth = os.Getenv("OAUTH_URL") + "/v1/authorization/" + target
		}
		requesters := requester.RequesterHandler()
		_, err := requesters.GET(uriOauth, tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token Not Valid",
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}
