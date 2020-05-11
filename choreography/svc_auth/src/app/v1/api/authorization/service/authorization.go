package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/authorization/entity"
	"github.com/sofyan48/svc_auth/src/middleware"
)

// AuthorizationService ...
type AuthorizationService struct {
	Middleware middleware.DefaultMiddlewareInterface
}

// AuthorizationServiceHandler ...
func AuthorizationServiceHandler() *AuthorizationService {
	return &AuthorizationService{
		Middleware: middleware.DefaultMiddlewareHandler(),
	}
}

// AuthorizationServiceInterface ...
type AuthorizationServiceInterface interface {
	ClientClaimsToken(ctx *gin.Context) (interface{}, error)
	UserClaimsToken(ctx *gin.Context) (interface{}, error)
}

// ClientClaimsToken ...
func (service *AuthorizationService) ClientClaimsToken(ctx *gin.Context) (interface{}, error) {
	tokenData, err := service.Middleware.GetSessionClaim(ctx)
	if err != nil {
		return nil, errors.New("Token Not Valid")
	}
	result := &entity.AuhtorizerClientResponse{}
	result.Expire = tokenData["exp"]
	result.Issuer = tokenData["iss"]
	result.Secret = tokenData["jti"]
	sessData := fmt.Sprintf("%s", tokenData["session"])
	sessions := &entity.AuhtorizerClientResponseSession{}
	json.Unmarshal([]byte(sessData), sessions)
	result.Session = sessions
	return result, nil
}

// UserClaimsToken ...
func (service *AuthorizationService) UserClaimsToken(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}
