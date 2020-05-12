package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_user/src/utils/requester"
)

// DefaultMiddleware types
type DefaultMiddleware struct {
	Requester requester.RequesterInterface
}

// DefaultMiddlewareHandler ...
func DefaultMiddlewareHandler() *DefaultMiddleware {
	return &DefaultMiddleware{
		Requester: requester.RequesterHandler(),
	}
}

// DefaultMiddlewareInterface ...
type DefaultMiddlewareInterface interface {
	GetToken(context *gin.Context) string
	AuthToken() gin.HandlerFunc
	CORSMiddleware() gin.HandlerFunc
}
