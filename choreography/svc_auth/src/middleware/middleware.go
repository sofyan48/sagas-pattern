package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/client/repository"
	"github.com/sofyan48/svc_auth/src/utils/redis"
)

// DefaultMiddleware types
type DefaultMiddleware struct {
	Client repository.ClientRepositoryInterface
	Redis  redis.RedisLibInterface
}

func DefaultMiddlewareHandler() *DefaultMiddleware {
	return &DefaultMiddleware{
		Client: repository.ClientRepositoryHandler(),
		Redis:  redis.RedisLibHandler(),
	}
}

type DefaultMiddlewareInterface interface {
	CORSMiddleware() gin.HandlerFunc
	GetToken(context *gin.Context) string
	AuthToken() gin.HandlerFunc
	GetSessionClaim(ctx *gin.Context) (map[string]interface{}, error)
}
