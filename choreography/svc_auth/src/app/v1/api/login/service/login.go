package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sofyan48/svc_auth/src/app/v1/api/login/entity"
	"github.com/sofyan48/svc_auth/src/middleware"
)

// LoginService ...
type LoginService struct {
	Middleware middleware.DefaultMiddlewareInterface
}

// LoginServiceHandler ...
func LoginServiceHandler() *LoginService {
	return &LoginService{
		Middleware: middleware.DefaultMiddlewareHandler(),
	}
}

// LoginServiceInterface ...
type LoginServiceInterface interface {
	PostLogin(payload *entity.LoginRequest, ctx *gin.Context) (*entity.LoginResponse, error)
}

// PostLogin ...
func (service *LoginService) PostLogin(payload *entity.LoginRequest, ctx *gin.Context) (*entity.LoginResponse, error) {
	claims, err := service.Middleware.GetSessionClaim(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(claims["exp"])
	fmt.Println(claims["jti"])
	fmt.Println(claims["session"])
	fmt.Println(claims["iss"])
	return nil, nil
}
